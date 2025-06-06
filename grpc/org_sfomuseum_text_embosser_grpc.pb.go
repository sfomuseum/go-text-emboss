// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.29.3
// source: grpc/org_sfomuseum_text_embosser.proto

package grpc

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

const (
	TextEmbosser_EmbossText_FullMethodName = "/org_sfomuseum_text_embosser.TextEmbosser/EmbossText"
)

// TextEmbosserClient is the client API for TextEmbosser service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextEmbosserClient interface {
	EmbossText(ctx context.Context, in *EmbossTextRequest, opts ...grpc.CallOption) (*EmbossTextResponse, error)
}

type textEmbosserClient struct {
	cc grpc.ClientConnInterface
}

func NewTextEmbosserClient(cc grpc.ClientConnInterface) TextEmbosserClient {
	return &textEmbosserClient{cc}
}

func (c *textEmbosserClient) EmbossText(ctx context.Context, in *EmbossTextRequest, opts ...grpc.CallOption) (*EmbossTextResponse, error) {
	out := new(EmbossTextResponse)
	err := c.cc.Invoke(ctx, TextEmbosser_EmbossText_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextEmbosserServer is the server API for TextEmbosser service.
// All implementations must embed UnimplementedTextEmbosserServer
// for forward compatibility
type TextEmbosserServer interface {
	EmbossText(context.Context, *EmbossTextRequest) (*EmbossTextResponse, error)
	mustEmbedUnimplementedTextEmbosserServer()
}

// UnimplementedTextEmbosserServer must be embedded to have forward compatible implementations.
type UnimplementedTextEmbosserServer struct {
}

func (UnimplementedTextEmbosserServer) EmbossText(context.Context, *EmbossTextRequest) (*EmbossTextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmbossText not implemented")
}
func (UnimplementedTextEmbosserServer) mustEmbedUnimplementedTextEmbosserServer() {}

// UnsafeTextEmbosserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextEmbosserServer will
// result in compilation errors.
type UnsafeTextEmbosserServer interface {
	mustEmbedUnimplementedTextEmbosserServer()
}

func RegisterTextEmbosserServer(s grpc.ServiceRegistrar, srv TextEmbosserServer) {
	s.RegisterService(&TextEmbosser_ServiceDesc, srv)
}

func _TextEmbosser_EmbossText_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmbossTextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextEmbosserServer).EmbossText(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TextEmbosser_EmbossText_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextEmbosserServer).EmbossText(ctx, req.(*EmbossTextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TextEmbosser_ServiceDesc is the grpc.ServiceDesc for TextEmbosser service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextEmbosser_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "org_sfomuseum_text_embosser.TextEmbosser",
	HandlerType: (*TextEmbosserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmbossText",
			Handler:    _TextEmbosser_EmbossText_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/org_sfomuseum_text_embosser.proto",
}
