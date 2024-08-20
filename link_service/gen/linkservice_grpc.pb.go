// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.2
// source: link_service/proto/linkservice.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LinkService_SaveLink_FullMethodName   = "/linkservice.LinkService/SaveLink"
	LinkService_GetLinks_FullMethodName   = "/linkservice.LinkService/GetLinks"
	LinkService_GetLink_FullMethodName    = "/linkservice.LinkService/GetLink"
	LinkService_DeleteLink_FullMethodName = "/linkservice.LinkService/DeleteLink"
)

// LinkServiceClient is the client API for LinkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkServiceClient interface {
	SaveLink(ctx context.Context, in *SaveLinkRequest, opts ...grpc.CallOption) (*SaveLinkResponse, error)
	GetLinks(ctx context.Context, in *GetLinksRequest, opts ...grpc.CallOption) (*GetLinksResponse, error)
	GetLink(ctx context.Context, in *GetLinkRequest, opts ...grpc.CallOption) (*GetLinkResponse, error)
	DeleteLink(ctx context.Context, in *DeleteLinkRequest, opts ...grpc.CallOption) (*DeleteLinkResponse, error)
}

type linkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkServiceClient(cc grpc.ClientConnInterface) LinkServiceClient {
	return &linkServiceClient{cc}
}

func (c *linkServiceClient) SaveLink(ctx context.Context, in *SaveLinkRequest, opts ...grpc.CallOption) (*SaveLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveLinkResponse)
	err := c.cc.Invoke(ctx, LinkService_SaveLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkServiceClient) GetLinks(ctx context.Context, in *GetLinksRequest, opts ...grpc.CallOption) (*GetLinksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLinksResponse)
	err := c.cc.Invoke(ctx, LinkService_GetLinks_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkServiceClient) GetLink(ctx context.Context, in *GetLinkRequest, opts ...grpc.CallOption) (*GetLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetLinkResponse)
	err := c.cc.Invoke(ctx, LinkService_GetLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkServiceClient) DeleteLink(ctx context.Context, in *DeleteLinkRequest, opts ...grpc.CallOption) (*DeleteLinkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteLinkResponse)
	err := c.cc.Invoke(ctx, LinkService_DeleteLink_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkServiceServer is the server API for LinkService service.
// All implementations must embed UnimplementedLinkServiceServer
// for forward compatibility.
type LinkServiceServer interface {
	SaveLink(context.Context, *SaveLinkRequest) (*SaveLinkResponse, error)
	GetLinks(context.Context, *GetLinksRequest) (*GetLinksResponse, error)
	GetLink(context.Context, *GetLinkRequest) (*GetLinkResponse, error)
	DeleteLink(context.Context, *DeleteLinkRequest) (*DeleteLinkResponse, error)
	mustEmbedUnimplementedLinkServiceServer()
}

// UnimplementedLinkServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLinkServiceServer struct{}

func (UnimplementedLinkServiceServer) SaveLink(context.Context, *SaveLinkRequest) (*SaveLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveLink not implemented")
}
func (UnimplementedLinkServiceServer) GetLinks(context.Context, *GetLinksRequest) (*GetLinksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinks not implemented")
}
func (UnimplementedLinkServiceServer) GetLink(context.Context, *GetLinkRequest) (*GetLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLink not implemented")
}
func (UnimplementedLinkServiceServer) DeleteLink(context.Context, *DeleteLinkRequest) (*DeleteLinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteLink not implemented")
}
func (UnimplementedLinkServiceServer) mustEmbedUnimplementedLinkServiceServer() {}
func (UnimplementedLinkServiceServer) testEmbeddedByValue()                     {}

// UnsafeLinkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkServiceServer will
// result in compilation errors.
type UnsafeLinkServiceServer interface {
	mustEmbedUnimplementedLinkServiceServer()
}

func RegisterLinkServiceServer(s grpc.ServiceRegistrar, srv LinkServiceServer) {
	// If the following call pancis, it indicates UnimplementedLinkServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LinkService_ServiceDesc, srv)
}

func _LinkService_SaveLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServiceServer).SaveLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkService_SaveLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServiceServer).SaveLink(ctx, req.(*SaveLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkService_GetLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServiceServer).GetLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkService_GetLinks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServiceServer).GetLinks(ctx, req.(*GetLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkService_GetLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServiceServer).GetLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkService_GetLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServiceServer).GetLink(ctx, req.(*GetLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkService_DeleteLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkServiceServer).DeleteLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkService_DeleteLink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkServiceServer).DeleteLink(ctx, req.(*DeleteLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkService_ServiceDesc is the grpc.ServiceDesc for LinkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "linkservice.LinkService",
	HandlerType: (*LinkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveLink",
			Handler:    _LinkService_SaveLink_Handler,
		},
		{
			MethodName: "GetLinks",
			Handler:    _LinkService_GetLinks_Handler,
		},
		{
			MethodName: "GetLink",
			Handler:    _LinkService_GetLink_Handler,
		},
		{
			MethodName: "DeleteLink",
			Handler:    _LinkService_DeleteLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "link_service/proto/linkservice.proto",
}
