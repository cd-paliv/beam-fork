/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.apache.beam.sdk.io.snowflake;

import java.io.IOException;
import java.io.StringReader;
import java.nio.charset.StandardCharsets;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.security.InvalidKeyException;
import java.security.KeyFactory;
import java.security.NoSuchAlgorithmException;
import java.security.PrivateKey;
import java.security.spec.InvalidKeySpecException;
import java.security.spec.PKCS8EncodedKeySpec;
import javax.crypto.EncryptedPrivateKeyInfo;
import javax.crypto.SecretKeyFactory;
import javax.crypto.spec.PBEKeySpec;
import org.apache.beam.vendor.guava.v32_1_2_jre.com.google.common.base.Strings;
import org.bouncycastle.util.encoders.Base64;
import org.bouncycastle.util.encoders.DecoderException;
import org.bouncycastle.util.io.pem.PemObject;
import org.bouncycastle.util.io.pem.PemReader;

public class KeyPairUtils {
  private static final String ENCRYPTED_PRIVATE_KEY = "ENCRYPTED PRIVATE KEY";
  private static final String UNENCRYPTED_PRIVATE_KEY = "PRIVATE KEY";

  private enum KeyEncryptionState {
    ENCRYPT,
    UNENCRYPTED,
    UNKNOWN
  }

  public static PrivateKey preparePrivateKey(String privateKey, String privateKeyPassphrase) {
    try {
      KeyFactory keyFactory = KeyFactory.getInstance("RSA");
      KeyEncryptionState encryptionState = guessKeyEncryptionState(privateKey);
      if (encryptionState == KeyEncryptionState.ENCRYPT
          && Strings.isNullOrEmpty(privateKeyPassphrase)) {
        throw new RuntimeException(
            "The private key is encrypted but no private key key passphrase has been provided.");
      }

      if (encryptionState == KeyEncryptionState.UNENCRYPTED
          && !Strings.isNullOrEmpty(privateKeyPassphrase)) {
        throw new RuntimeException(
            "The private key is unencrypted but private key key passphrase has been provided.");
      }

      byte[] decoded;

      if (encryptionState == KeyEncryptionState.UNKNOWN) {
        decoded = Base64.decode(privateKey);
      } else {
        PemReader pr = new PemReader(new StringReader(privateKey));
        PemObject pemObject = pr.readPemObject();
        decoded = pemObject.getContent();
        pr.close();
      }

      if (Strings.isNullOrEmpty(privateKeyPassphrase)) {
        // unencrypted private key file
        PKCS8EncodedKeySpec encodedKeySpec = new PKCS8EncodedKeySpec(decoded);
        return keyFactory.generatePrivate(encodedKeySpec);
      } else {
        // encrypted private key file
        EncryptedPrivateKeyInfo pkInfo = new EncryptedPrivateKeyInfo(decoded);
        PBEKeySpec keySpec = new PBEKeySpec(privateKeyPassphrase.toCharArray());
        SecretKeyFactory pbeKeyFactory = SecretKeyFactory.getInstance(pkInfo.getAlgName());
        PKCS8EncodedKeySpec encodedKeySpec =
            pkInfo.getKeySpec(pbeKeyFactory.generateSecret(keySpec));
        return keyFactory.generatePrivate(encodedKeySpec);
      }
    } catch (NoSuchAlgorithmException e) {
      throw new RuntimeException(
          "Private key encryption algorithm not supported. This may mean that the private key was generated by OpenSSL 1.1.1g or newer "
              + "which uses an encryption algorithm by default which has compatibility issues in some JVM environments. "
              + "For details, see: "
              + "https://community.snowflake.com/s/article/Private-key-provided-is-invalid-or-not-supported-rsa-key-p8--data-isn-t-an-object-ID"
              + " "
              + e.getMessage());
    } catch (InvalidKeySpecException
        | IOException
        | IllegalArgumentException
        | NullPointerException
        | InvalidKeyException
        | DecoderException e) {
      throw new RuntimeException("Can't create private key: " + e.getMessage(), e);
    }
  }

  /**
   * Tries to determine whether the private key is encrypted or not based on the file headers.
   *
   * <p>If this is not possible (e.g. there are no headers), returns {@link
   * KeyEncryptionState#UNKNOWN}
   */
  private static KeyEncryptionState guessKeyEncryptionState(String privateKey) {
    PemReader pr = new PemReader(new StringReader(privateKey));
    try {
      PemObject pemObject = pr.readPemObject();
      if (pemObject == null) {
        // If it is not a PEM file then it is not possible to determine the encryption state
        return KeyEncryptionState.UNKNOWN;
      }
      if (ENCRYPTED_PRIVATE_KEY.equals(pemObject.getType())) {
        return KeyEncryptionState.ENCRYPT;
      } else if (UNENCRYPTED_PRIVATE_KEY.equals(pemObject.getType())) {
        return KeyEncryptionState.UNENCRYPTED;
      } else {
        throw new RuntimeException(
            "Invalid type of PEM file: "
                + pemObject.getType()
                + ". Supported types: "
                + ENCRYPTED_PRIVATE_KEY
                + ", "
                + UNENCRYPTED_PRIVATE_KEY);
      }
    } catch (IOException e) {
      throw new RuntimeException("Can't read parse private key");
    }
  }

  public static String readPrivateKeyFile(String privateKeyPath) {
    try {
      byte[] keyBytes = Files.readAllBytes(Paths.get(privateKeyPath));
      return new String(keyBytes, StandardCharsets.UTF_8);
    } catch (IOException e) {
      throw new RuntimeException("Can't read private key from provided path");
    }
  }
}