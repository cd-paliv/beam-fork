//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//
// Protocol Buffers describing the Expansion API, an api for expanding
// transforms in a remote SDK.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: org/apache/beam/model/job_management/v1/beam_expansion_api.proto

package jobmanagement_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ExpansionService_Expand_FullMethodName                  = "/org.apache.beam.model.expansion.v1.ExpansionService/Expand"
	ExpansionService_DiscoverSchemaTransform_FullMethodName = "/org.apache.beam.model.expansion.v1.ExpansionService/DiscoverSchemaTransform"
)

// ExpansionServiceClient is the client API for ExpansionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Job Service for constructing pipelines
type ExpansionServiceClient interface {
	Expand(ctx context.Context, in *ExpansionRequest, opts ...grpc.CallOption) (*ExpansionResponse, error)
	// A RPC to discover already registered SchemaTransformProviders.
	// See https://s.apache.org/easy-multi-language for more details.
	DiscoverSchemaTransform(ctx context.Context, in *DiscoverSchemaTransformRequest, opts ...grpc.CallOption) (*DiscoverSchemaTransformResponse, error)
}

type expansionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExpansionServiceClient(cc grpc.ClientConnInterface) ExpansionServiceClient {
	return &expansionServiceClient{cc}
}

func (c *expansionServiceClient) Expand(ctx context.Context, in *ExpansionRequest, opts ...grpc.CallOption) (*ExpansionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExpansionResponse)
	err := c.cc.Invoke(ctx, ExpansionService_Expand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *expansionServiceClient) DiscoverSchemaTransform(ctx context.Context, in *DiscoverSchemaTransformRequest, opts ...grpc.CallOption) (*DiscoverSchemaTransformResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DiscoverSchemaTransformResponse)
	err := c.cc.Invoke(ctx, ExpansionService_DiscoverSchemaTransform_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExpansionServiceServer is the server API for ExpansionService service.
// All implementations must embed UnimplementedExpansionServiceServer
// for forward compatibility
//
// Job Service for constructing pipelines
type ExpansionServiceServer interface {
	Expand(context.Context, *ExpansionRequest) (*ExpansionResponse, error)
	// A RPC to discover already registered SchemaTransformProviders.
	// See https://s.apache.org/easy-multi-language for more details.
	DiscoverSchemaTransform(context.Context, *DiscoverSchemaTransformRequest) (*DiscoverSchemaTransformResponse, error)
	mustEmbedUnimplementedExpansionServiceServer()
}

// UnimplementedExpansionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedExpansionServiceServer struct {
}

func (UnimplementedExpansionServiceServer) Expand(context.Context, *ExpansionRequest) (*ExpansionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Expand not implemented")
}
func (UnimplementedExpansionServiceServer) DiscoverSchemaTransform(context.Context, *DiscoverSchemaTransformRequest) (*DiscoverSchemaTransformResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DiscoverSchemaTransform not implemented")
}
func (UnimplementedExpansionServiceServer) mustEmbedUnimplementedExpansionServiceServer() {}

// UnsafeExpansionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExpansionServiceServer will
// result in compilation errors.
type UnsafeExpansionServiceServer interface {
	mustEmbedUnimplementedExpansionServiceServer()
}

func RegisterExpansionServiceServer(s grpc.ServiceRegistrar, srv ExpansionServiceServer) {
	s.RegisterService(&ExpansionService_ServiceDesc, srv)
}

func _ExpansionService_Expand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpansionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpansionServiceServer).Expand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExpansionService_Expand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpansionServiceServer).Expand(ctx, req.(*ExpansionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExpansionService_DiscoverSchemaTransform_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoverSchemaTransformRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExpansionServiceServer).DiscoverSchemaTransform(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExpansionService_DiscoverSchemaTransform_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExpansionServiceServer).DiscoverSchemaTransform(ctx, req.(*DiscoverSchemaTransformRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExpansionService_ServiceDesc is the grpc.ServiceDesc for ExpansionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExpansionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.apache.beam.model.expansion.v1.ExpansionService",
	HandlerType: (*ExpansionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Expand",
			Handler:    _ExpansionService_Expand_Handler,
		},
		{
			MethodName: "DiscoverSchemaTransform",
			Handler:    _ExpansionService_DiscoverSchemaTransform_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "org/apache/beam/model/job_management/v1/beam_expansion_api.proto",
}
