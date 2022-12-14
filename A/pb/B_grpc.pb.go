// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: B.proto

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

// PostServicesClient is the client API for PostServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PostServicesClient interface {
	InsertPosts(ctx context.Context, in *InsertPostRequest, opts ...grpc.CallOption) (*InsertPostRequest, error)
	SearchPosts(ctx context.Context, in *SearchPostRequest, opts ...grpc.CallOption) (*SearchPostResponse, error)
	FindAllPosts(ctx context.Context, in *FindAllRequest, opts ...grpc.CallOption) (*FindAllResponse, error)
	UpdatePosts(ctx context.Context, in *UpdateNew, opts ...grpc.CallOption) (*UpdatePostResponse, error)
	DeletePosts(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error)
}

type postServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewPostServicesClient(cc grpc.ClientConnInterface) PostServicesClient {
	return &postServicesClient{cc}
}

func (c *postServicesClient) InsertPosts(ctx context.Context, in *InsertPostRequest, opts ...grpc.CallOption) (*InsertPostRequest, error) {
	out := new(InsertPostRequest)
	err := c.cc.Invoke(ctx, "/Hello.PostServices/InsertPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServicesClient) SearchPosts(ctx context.Context, in *SearchPostRequest, opts ...grpc.CallOption) (*SearchPostResponse, error) {
	out := new(SearchPostResponse)
	err := c.cc.Invoke(ctx, "/Hello.PostServices/SearchPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServicesClient) FindAllPosts(ctx context.Context, in *FindAllRequest, opts ...grpc.CallOption) (*FindAllResponse, error) {
	out := new(FindAllResponse)
	err := c.cc.Invoke(ctx, "/Hello.PostServices/FindAllPosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServicesClient) UpdatePosts(ctx context.Context, in *UpdateNew, opts ...grpc.CallOption) (*UpdatePostResponse, error) {
	out := new(UpdatePostResponse)
	err := c.cc.Invoke(ctx, "/Hello.PostServices/UpdatePosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postServicesClient) DeletePosts(ctx context.Context, in *DeletePostRequest, opts ...grpc.CallOption) (*DeletePostResponse, error) {
	out := new(DeletePostResponse)
	err := c.cc.Invoke(ctx, "/Hello.PostServices/DeletePosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostServicesServer is the server API for PostServices service.
// All implementations must embed UnimplementedPostServicesServer
// for forward compatibility
type PostServicesServer interface {
	InsertPosts(context.Context, *InsertPostRequest) (*InsertPostRequest, error)
	SearchPosts(context.Context, *SearchPostRequest) (*SearchPostResponse, error)
	FindAllPosts(context.Context, *FindAllRequest) (*FindAllResponse, error)
	UpdatePosts(context.Context, *UpdateNew) (*UpdatePostResponse, error)
	DeletePosts(context.Context, *DeletePostRequest) (*DeletePostResponse, error)
	mustEmbedUnimplementedPostServicesServer()
}

// UnimplementedPostServicesServer must be embedded to have forward compatible implementations.
type UnimplementedPostServicesServer struct {
}

func (UnimplementedPostServicesServer) InsertPosts(context.Context, *InsertPostRequest) (*InsertPostRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertPosts not implemented")
}
func (UnimplementedPostServicesServer) SearchPosts(context.Context, *SearchPostRequest) (*SearchPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPosts not implemented")
}
func (UnimplementedPostServicesServer) FindAllPosts(context.Context, *FindAllRequest) (*FindAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllPosts not implemented")
}
func (UnimplementedPostServicesServer) UpdatePosts(context.Context, *UpdateNew) (*UpdatePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePosts not implemented")
}
func (UnimplementedPostServicesServer) DeletePosts(context.Context, *DeletePostRequest) (*DeletePostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePosts not implemented")
}
func (UnimplementedPostServicesServer) mustEmbedUnimplementedPostServicesServer() {}

// UnsafePostServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PostServicesServer will
// result in compilation errors.
type UnsafePostServicesServer interface {
	mustEmbedUnimplementedPostServicesServer()
}

func RegisterPostServicesServer(s grpc.ServiceRegistrar, srv PostServicesServer) {
	s.RegisterService(&PostServices_ServiceDesc, srv)
}

func _PostServices_InsertPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServicesServer).InsertPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hello.PostServices/InsertPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServicesServer).InsertPosts(ctx, req.(*InsertPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostServices_SearchPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServicesServer).SearchPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hello.PostServices/SearchPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServicesServer).SearchPosts(ctx, req.(*SearchPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostServices_FindAllPosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServicesServer).FindAllPosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hello.PostServices/FindAllPosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServicesServer).FindAllPosts(ctx, req.(*FindAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostServices_UpdatePosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServicesServer).UpdatePosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hello.PostServices/UpdatePosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServicesServer).UpdatePosts(ctx, req.(*UpdateNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostServices_DeletePosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostServicesServer).DeletePosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Hello.PostServices/DeletePosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostServicesServer).DeletePosts(ctx, req.(*DeletePostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PostServices_ServiceDesc is the grpc.ServiceDesc for PostServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PostServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Hello.PostServices",
	HandlerType: (*PostServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InsertPosts",
			Handler:    _PostServices_InsertPosts_Handler,
		},
		{
			MethodName: "SearchPosts",
			Handler:    _PostServices_SearchPosts_Handler,
		},
		{
			MethodName: "FindAllPosts",
			Handler:    _PostServices_FindAllPosts_Handler,
		},
		{
			MethodName: "UpdatePosts",
			Handler:    _PostServices_UpdatePosts_Handler,
		},
		{
			MethodName: "DeletePosts",
			Handler:    _PostServices_DeletePosts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "B.proto",
}
