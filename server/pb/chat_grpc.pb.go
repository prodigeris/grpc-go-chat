// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: chat.proto

package pb

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

// TranscriptClient is the client API for Transcript service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TranscriptClient interface {
	GetChat(ctx context.Context, in *GetMessageListRequest, opts ...grpc.CallOption) (*GetMessageListResponse, error)
}

type transcriptClient struct {
	cc grpc.ClientConnInterface
}

func NewTranscriptClient(cc grpc.ClientConnInterface) TranscriptClient {
	return &transcriptClient{cc}
}

func (c *transcriptClient) GetChat(ctx context.Context, in *GetMessageListRequest, opts ...grpc.CallOption) (*GetMessageListResponse, error) {
	out := new(GetMessageListResponse)
	err := c.cc.Invoke(ctx, "/Transcript/GetChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranscriptServer is the server API for Transcript service.
// All implementations must embed UnimplementedTranscriptServer
// for forward compatibility
type TranscriptServer interface {
	GetChat(context.Context, *GetMessageListRequest) (*GetMessageListResponse, error)
	mustEmbedUnimplementedTranscriptServer()
}

// UnimplementedTranscriptServer must be embedded to have forward compatible implementations.
type UnimplementedTranscriptServer struct {
}

func (UnimplementedTranscriptServer) GetChat(context.Context, *GetMessageListRequest) (*GetMessageListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChat not implemented")
}
func (UnimplementedTranscriptServer) mustEmbedUnimplementedTranscriptServer() {}

// UnsafeTranscriptServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TranscriptServer will
// result in compilation errors.
type UnsafeTranscriptServer interface {
	mustEmbedUnimplementedTranscriptServer()
}

func RegisterTranscriptServer(s grpc.ServiceRegistrar, srv TranscriptServer) {
	s.RegisterService(&Transcript_ServiceDesc, srv)
}

func _Transcript_GetChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranscriptServer).GetChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Transcript/GetChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranscriptServer).GetChat(ctx, req.(*GetMessageListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Transcript_ServiceDesc is the grpc.ServiceDesc for Transcript service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Transcript_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Transcript",
	HandlerType: (*TranscriptServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChat",
			Handler:    _Transcript_GetChat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chat.proto",
}
