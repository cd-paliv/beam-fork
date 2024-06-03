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
// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.1.0
// - protoc             v3.12.4
// source: org/apache/beam/model/pipeline/v1/beam_runner_api.proto

package pipeline_v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TestStreamServiceClient is the client API for TestStreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestStreamServiceClient interface {
	// A TestStream will request for events using this RPC.
	Events(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (TestStreamService_EventsClient, error)
}

type testStreamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestStreamServiceClient(cc grpc.ClientConnInterface) TestStreamServiceClient {
	return &testStreamServiceClient{cc}
}

func (c *testStreamServiceClient) Events(ctx context.Context, in *EventsRequest, opts ...grpc.CallOption) (TestStreamService_EventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &TestStreamService_ServiceDesc.Streams[0], "/org.apache.beam.model.pipeline.v1.TestStreamService/Events", opts...)
	if err != nil {
		return nil, err
	}
	x := &testStreamServiceEventsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TestStreamService_EventsClient interface {
	Recv() (*TestStreamPayload_Event, error)
	grpc.ClientStream
}

type testStreamServiceEventsClient struct {
	grpc.ClientStream
}

func (x *testStreamServiceEventsClient) Recv() (*TestStreamPayload_Event, error) {
	m := new(TestStreamPayload_Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestStreamServiceServer is the server API for TestStreamService service.
// All implementations must embed UnimplementedTestStreamServiceServer
// for forward compatibility
type TestStreamServiceServer interface {
	// A TestStream will request for events using this RPC.
	Events(*EventsRequest, TestStreamService_EventsServer) error
	mustEmbedUnimplementedTestStreamServiceServer()
}

// UnimplementedTestStreamServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTestStreamServiceServer struct {
}

func (UnimplementedTestStreamServiceServer) Events(*EventsRequest, TestStreamService_EventsServer) error {
	return status.Errorf(codes.Unimplemented, "method Events not implemented")
}
func (UnimplementedTestStreamServiceServer) mustEmbedUnimplementedTestStreamServiceServer() {}

// UnsafeTestStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestStreamServiceServer will
// result in compilation errors.
type UnsafeTestStreamServiceServer interface {
	mustEmbedUnimplementedTestStreamServiceServer()
}

func RegisterTestStreamServiceServer(s grpc.ServiceRegistrar, srv TestStreamServiceServer) {
	s.RegisterService(&TestStreamService_ServiceDesc, srv)
}

func _TestStreamService_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestStreamServiceServer).Events(m, &testStreamServiceEventsServer{stream})
}

type TestStreamService_EventsServer interface {
	Send(*TestStreamPayload_Event) error
	grpc.ServerStream
}

type testStreamServiceEventsServer struct {
	grpc.ServerStream
}

func (x *testStreamServiceEventsServer) Send(m *TestStreamPayload_Event) error {
	return x.ServerStream.SendMsg(m)
}

// TestStreamService_ServiceDesc is the grpc.ServiceDesc for TestStreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestStreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org.apache.beam.model.pipeline.v1.TestStreamService",
	HandlerType: (*TestStreamServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Events",
			Handler:       _TestStreamService_Events_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "org/apache/beam/model/pipeline/v1/beam_runner_api.proto",
}
