// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.2
// source: service/task_app.proto

package service

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
	TaskProto_Create_FullMethodName  = "/service.TaskProto/Create"
	TaskProto_List_FullMethodName    = "/service.TaskProto/List"
	TaskProto_Details_FullMethodName = "/service.TaskProto/Details"
	TaskProto_Update_FullMethodName  = "/service.TaskProto/Update"
)

// TaskProtoClient is the client API for TaskProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TaskProtoClient interface {
	Create(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error)
	List(ctx context.Context, in *ListTaskRequest, opts ...grpc.CallOption) (*ListTaskResponse, error)
	Details(ctx context.Context, in *DetailsTaskRequest, opts ...grpc.CallOption) (*DetailsTaskResponse, error)
	Update(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error)
}

type taskProtoClient struct {
	cc grpc.ClientConnInterface
}

func NewTaskProtoClient(cc grpc.ClientConnInterface) TaskProtoClient {
	return &taskProtoClient{cc}
}

func (c *taskProtoClient) Create(ctx context.Context, in *CreateTaskRequest, opts ...grpc.CallOption) (*CreateTaskResponse, error) {
	out := new(CreateTaskResponse)
	err := c.cc.Invoke(ctx, TaskProto_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskProtoClient) List(ctx context.Context, in *ListTaskRequest, opts ...grpc.CallOption) (*ListTaskResponse, error) {
	out := new(ListTaskResponse)
	err := c.cc.Invoke(ctx, TaskProto_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskProtoClient) Details(ctx context.Context, in *DetailsTaskRequest, opts ...grpc.CallOption) (*DetailsTaskResponse, error) {
	out := new(DetailsTaskResponse)
	err := c.cc.Invoke(ctx, TaskProto_Details_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *taskProtoClient) Update(ctx context.Context, in *UpdateTaskRequest, opts ...grpc.CallOption) (*UpdateTaskResponse, error) {
	out := new(UpdateTaskResponse)
	err := c.cc.Invoke(ctx, TaskProto_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TaskProtoServer is the server API for TaskProto service.
// All implementations must embed UnimplementedTaskProtoServer
// for forward compatibility
type TaskProtoServer interface {
	Create(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error)
	List(context.Context, *ListTaskRequest) (*ListTaskResponse, error)
	Details(context.Context, *DetailsTaskRequest) (*DetailsTaskResponse, error)
	Update(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error)
	mustEmbedUnimplementedTaskProtoServer()
}

// UnimplementedTaskProtoServer must be embedded to have forward compatible implementations.
type UnimplementedTaskProtoServer struct {
}

func (UnimplementedTaskProtoServer) Create(context.Context, *CreateTaskRequest) (*CreateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTaskProtoServer) List(context.Context, *ListTaskRequest) (*ListTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTaskProtoServer) Details(context.Context, *DetailsTaskRequest) (*DetailsTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Details not implemented")
}
func (UnimplementedTaskProtoServer) Update(context.Context, *UpdateTaskRequest) (*UpdateTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTaskProtoServer) mustEmbedUnimplementedTaskProtoServer() {}

// UnsafeTaskProtoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TaskProtoServer will
// result in compilation errors.
type UnsafeTaskProtoServer interface {
	mustEmbedUnimplementedTaskProtoServer()
}

func RegisterTaskProtoServer(s grpc.ServiceRegistrar, srv TaskProtoServer) {
	s.RegisterService(&TaskProto_ServiceDesc, srv)
}

func _TaskProto_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskProtoServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskProto_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskProtoServer).Create(ctx, req.(*CreateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskProto_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskProtoServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskProto_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskProtoServer).List(ctx, req.(*ListTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskProto_Details_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetailsTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskProtoServer).Details(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskProto_Details_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskProtoServer).Details(ctx, req.(*DetailsTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TaskProto_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TaskProtoServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TaskProto_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TaskProtoServer).Update(ctx, req.(*UpdateTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TaskProto_ServiceDesc is the grpc.ServiceDesc for TaskProto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TaskProto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.TaskProto",
	HandlerType: (*TaskProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TaskProto_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TaskProto_List_Handler,
		},
		{
			MethodName: "Details",
			Handler:    _TaskProto_Details_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TaskProto_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/task_app.proto",
}
