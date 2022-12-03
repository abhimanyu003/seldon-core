/*
Copyright 2022 Seldon Technologies Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.10
// source: mlops/chainer/chainer.proto

package chainer

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

// ChainerClient is the client API for Chainer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChainerClient interface {
	SubscribePipelineUpdates(ctx context.Context, in *PipelineSubscriptionRequest, opts ...grpc.CallOption) (Chainer_SubscribePipelineUpdatesClient, error)
	PipelineUpdateEvent(ctx context.Context, in *PipelineUpdateStatusMessage, opts ...grpc.CallOption) (*PipelineUpdateStatusResponse, error)
}

type chainerClient struct {
	cc grpc.ClientConnInterface
}

func NewChainerClient(cc grpc.ClientConnInterface) ChainerClient {
	return &chainerClient{cc}
}

func (c *chainerClient) SubscribePipelineUpdates(ctx context.Context, in *PipelineSubscriptionRequest, opts ...grpc.CallOption) (Chainer_SubscribePipelineUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chainer_ServiceDesc.Streams[0], "/seldon.mlops.chainer.Chainer/SubscribePipelineUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &chainerSubscribePipelineUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chainer_SubscribePipelineUpdatesClient interface {
	Recv() (*PipelineUpdateMessage, error)
	grpc.ClientStream
}

type chainerSubscribePipelineUpdatesClient struct {
	grpc.ClientStream
}

func (x *chainerSubscribePipelineUpdatesClient) Recv() (*PipelineUpdateMessage, error) {
	m := new(PipelineUpdateMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chainerClient) PipelineUpdateEvent(ctx context.Context, in *PipelineUpdateStatusMessage, opts ...grpc.CallOption) (*PipelineUpdateStatusResponse, error) {
	out := new(PipelineUpdateStatusResponse)
	err := c.cc.Invoke(ctx, "/seldon.mlops.chainer.Chainer/PipelineUpdateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChainerServer is the server API for Chainer service.
// All implementations must embed UnimplementedChainerServer
// for forward compatibility
type ChainerServer interface {
	SubscribePipelineUpdates(*PipelineSubscriptionRequest, Chainer_SubscribePipelineUpdatesServer) error
	PipelineUpdateEvent(context.Context, *PipelineUpdateStatusMessage) (*PipelineUpdateStatusResponse, error)
	mustEmbedUnimplementedChainerServer()
}

// UnimplementedChainerServer must be embedded to have forward compatible implementations.
type UnimplementedChainerServer struct {
}

func (UnimplementedChainerServer) SubscribePipelineUpdates(*PipelineSubscriptionRequest, Chainer_SubscribePipelineUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribePipelineUpdates not implemented")
}
func (UnimplementedChainerServer) PipelineUpdateEvent(context.Context, *PipelineUpdateStatusMessage) (*PipelineUpdateStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PipelineUpdateEvent not implemented")
}
func (UnimplementedChainerServer) mustEmbedUnimplementedChainerServer() {}

// UnsafeChainerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChainerServer will
// result in compilation errors.
type UnsafeChainerServer interface {
	mustEmbedUnimplementedChainerServer()
}

func RegisterChainerServer(s grpc.ServiceRegistrar, srv ChainerServer) {
	s.RegisterService(&Chainer_ServiceDesc, srv)
}

func _Chainer_SubscribePipelineUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PipelineSubscriptionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChainerServer).SubscribePipelineUpdates(m, &chainerSubscribePipelineUpdatesServer{stream})
}

type Chainer_SubscribePipelineUpdatesServer interface {
	Send(*PipelineUpdateMessage) error
	grpc.ServerStream
}

type chainerSubscribePipelineUpdatesServer struct {
	grpc.ServerStream
}

func (x *chainerSubscribePipelineUpdatesServer) Send(m *PipelineUpdateMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _Chainer_PipelineUpdateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PipelineUpdateStatusMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChainerServer).PipelineUpdateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/seldon.mlops.chainer.Chainer/PipelineUpdateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChainerServer).PipelineUpdateEvent(ctx, req.(*PipelineUpdateStatusMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Chainer_ServiceDesc is the grpc.ServiceDesc for Chainer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chainer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "seldon.mlops.chainer.Chainer",
	HandlerType: (*ChainerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PipelineUpdateEvent",
			Handler:    _Chainer_PipelineUpdateEvent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribePipelineUpdates",
			Handler:       _Chainer_SubscribePipelineUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "mlops/chainer/chainer.proto",
}
