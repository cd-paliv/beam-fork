// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// boot is the boot code for the Java SDK harness container. It is responsible
// for retrieving staged files and invoking the JVM correctly.
package main

import (
	"context"
	"reflect"
	"testing"

	"github.com/cd-paliv/beam-fork/sdks/v3/go/container/tools"
)

func TestBuildOptionsEmpty(t *testing.T) {
	ctx, logger := context.Background(), &tools.Logger{}
	dir := "test/empty"
	metaOptions, err := LoadMetaOptions(ctx, logger, dir)
	if err != nil {
		t.Fatalf("Got error %v running LoadMetaOptions", err)
	}
	if metaOptions != nil {
		t.Fatalf("LoadMetaOptions(%v) = %v, want nil", dir, metaOptions)
	}

	javaOptions := BuildOptions(ctx, logger, metaOptions)
	if len(javaOptions.JavaArguments) != 0 || len(javaOptions.Classpath) != 0 || len(javaOptions.Properties) != 0 {
		t.Errorf("BuildOptions(%v) = %v, want nil", metaOptions, javaOptions)
	}
}

func TestBuildOptionsDisabled(t *testing.T) {
	ctx, logger := context.Background(), &tools.Logger{}
	metaOptions, err := LoadMetaOptions(ctx, logger, "test/disabled")
	if err != nil {
		t.Fatalf("Got error %v running LoadMetaOptions", err)
	}

	javaOptions := BuildOptions(ctx, logger, metaOptions)
	if len(javaOptions.JavaArguments) != 0 || len(javaOptions.Classpath) != 0 || len(javaOptions.Properties) != 0 {
		t.Errorf("BuildOptions(%v) = %v, want nil", metaOptions, javaOptions)
	}
}

func TestBuildOptions(t *testing.T) {
	ctx, logger := context.Background(), &tools.Logger{}
	metaOptions, err := LoadMetaOptions(ctx, logger, "test/priority")
	if err != nil {
		t.Fatalf("Got error %v running LoadMetaOptions", err)
	}

	javaOptions := BuildOptions(ctx, logger, metaOptions)
	wantJavaArguments := []string{"java_args=low", "java_args=high"}
	wantClasspath := []string{"classpath_high", "classpath_low"}
	wantProperties := map[string]string{
		"priority": "high",
	}
	if !reflect.DeepEqual(javaOptions.JavaArguments, wantJavaArguments) {
		t.Errorf("BuildOptions(%v).JavaArguments = %v, want %v", metaOptions, javaOptions.JavaArguments, wantJavaArguments)
	}
	if !reflect.DeepEqual(javaOptions.Classpath, wantClasspath) {
		t.Errorf("BuildOptions(%v).Classpath = %v, want %v", metaOptions, javaOptions.Classpath, wantClasspath)
	}
	if !reflect.DeepEqual(javaOptions.Properties, wantProperties) {
		t.Errorf("BuildOptions(%v).JavaProperties = %v, want %v", metaOptions, javaOptions.Properties, wantProperties)
	}
}
