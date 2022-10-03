// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: object.proto

package object

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

// ObjectStoreClient is the client API for ObjectStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ObjectStoreClient interface {
	Read(ctx context.Context, in *ReadObjectRequest, opts ...grpc.CallOption) (*ReadObjectResponse, error)
	BatchRead(ctx context.Context, in *BatchReadObjectRequest, opts ...grpc.CallOption) (*BatchReadObjectResponse, error)
	Write(ctx context.Context, in *WriteObjectRequest, opts ...grpc.CallOption) (*WriteObjectResponse, error)
	Delete(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error)
	History(ctx context.Context, in *ObjectHistoryRequest, opts ...grpc.CallOption) (*ObjectHistoryResponse, error)
	Search(ctx context.Context, in *ObjectSearchRequest, opts ...grpc.CallOption) (*ObjectSearchResponse, error)
}

type objectStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewObjectStoreClient(cc grpc.ClientConnInterface) ObjectStoreClient {
	return &objectStoreClient{cc}
}

func (c *objectStoreClient) Read(ctx context.Context, in *ReadObjectRequest, opts ...grpc.CallOption) (*ReadObjectResponse, error) {
	out := new(ReadObjectResponse)
	err := c.cc.Invoke(ctx, "/object.ObjectStore/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) BatchRead(ctx context.Context, in *BatchReadObjectRequest, opts ...grpc.CallOption) (*BatchReadObjectResponse, error) {
	out := new(BatchReadObjectResponse)
	err := c.cc.Invoke(ctx, "/object.ObjectStore/BatchRead", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) Write(ctx context.Context, in *WriteObjectRequest, opts ...grpc.CallOption) (*WriteObjectResponse, error) {
	out := new(WriteObjectResponse)
	err := c.cc.Invoke(ctx, "/object.ObjectStore/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) Delete(ctx context.Context, in *DeleteObjectRequest, opts ...grpc.CallOption) (*DeleteObjectResponse, error) {
	out := new(DeleteObjectResponse)
	err := c.cc.Invoke(ctx, "/object.ObjectStore/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) History(ctx context.Context, in *ObjectHistoryRequest, opts ...grpc.CallOption) (*ObjectHistoryResponse, error) {
	out := new(ObjectHistoryResponse)
	err := c.cc.Invoke(ctx, "/object.ObjectStore/History", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *objectStoreClient) Search(ctx context.Context, in *ObjectSearchRequest, opts ...grpc.CallOption) (*ObjectSearchResponse, error) {
	out := new(ObjectSearchResponse)
	err := c.cc.Invoke(ctx, "/object.ObjectStore/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ObjectStoreServer is the server API for ObjectStore service.
// All implementations should embed UnimplementedObjectStoreServer
// for forward compatibility
type ObjectStoreServer interface {
	Read(context.Context, *ReadObjectRequest) (*ReadObjectResponse, error)
	BatchRead(context.Context, *BatchReadObjectRequest) (*BatchReadObjectResponse, error)
	Write(context.Context, *WriteObjectRequest) (*WriteObjectResponse, error)
	Delete(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error)
	History(context.Context, *ObjectHistoryRequest) (*ObjectHistoryResponse, error)
	Search(context.Context, *ObjectSearchRequest) (*ObjectSearchResponse, error)
}

// UnimplementedObjectStoreServer should be embedded to have forward compatible implementations.
type UnimplementedObjectStoreServer struct {
}

func (UnimplementedObjectStoreServer) Read(context.Context, *ReadObjectRequest) (*ReadObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedObjectStoreServer) BatchRead(context.Context, *BatchReadObjectRequest) (*BatchReadObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchRead not implemented")
}
func (UnimplementedObjectStoreServer) Write(context.Context, *WriteObjectRequest) (*WriteObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedObjectStoreServer) Delete(context.Context, *DeleteObjectRequest) (*DeleteObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedObjectStoreServer) History(context.Context, *ObjectHistoryRequest) (*ObjectHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method History not implemented")
}
func (UnimplementedObjectStoreServer) Search(context.Context, *ObjectSearchRequest) (*ObjectSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

// UnsafeObjectStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ObjectStoreServer will
// result in compilation errors.
type UnsafeObjectStoreServer interface {
	mustEmbedUnimplementedObjectStoreServer()
}

func RegisterObjectStoreServer(s grpc.ServiceRegistrar, srv ObjectStoreServer) {
	s.RegisterService(&ObjectStore_ServiceDesc, srv)
}

func _ObjectStore_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object.ObjectStore/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).Read(ctx, req.(*ReadObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_BatchRead_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchReadObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).BatchRead(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object.ObjectStore/BatchRead",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).BatchRead(ctx, req.(*BatchReadObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object.ObjectStore/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).Write(ctx, req.(*WriteObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteObjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object.ObjectStore/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).Delete(ctx, req.(*DeleteObjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_History_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).History(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object.ObjectStore/History",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).History(ctx, req.(*ObjectHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ObjectStore_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ObjectSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ObjectStoreServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/object.ObjectStore/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ObjectStoreServer).Search(ctx, req.(*ObjectSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ObjectStore_ServiceDesc is the grpc.ServiceDesc for ObjectStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ObjectStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "object.ObjectStore",
	HandlerType: (*ObjectStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _ObjectStore_Read_Handler,
		},
		{
			MethodName: "BatchRead",
			Handler:    _ObjectStore_BatchRead_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _ObjectStore_Write_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ObjectStore_Delete_Handler,
		},
		{
			MethodName: "History",
			Handler:    _ObjectStore_History_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _ObjectStore_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "object.proto",
}
