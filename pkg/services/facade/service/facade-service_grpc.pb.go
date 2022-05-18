// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// FacadeServiceClient is the client API for FacadeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FacadeServiceClient interface {
	Signup(ctx context.Context, in *SignupRequset, opts ...grpc.CallOption) (*SignupResponse, error)
	AllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAlllUsersResponse, error)
	// Create new task
	CreateTask(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Read task
	GetTaskById(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	// Update task
	UpdateTask(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete task
	DeleteTask(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Read all tasks
	ReadAllTasks(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error)
}

type facadeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFacadeServiceClient(cc grpc.ClientConnInterface) FacadeServiceClient {
	return &facadeServiceClient{cc}
}

func (c *facadeServiceClient) Signup(ctx context.Context, in *SignupRequset, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, "/facade.FacadeService/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *facadeServiceClient) AllUsers(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*GetAlllUsersResponse, error) {
	out := new(GetAlllUsersResponse)
	err := c.cc.Invoke(ctx, "/facade.FacadeService/AllUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *facadeServiceClient) CreateTask(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/facade.FacadeService/CreateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *facadeServiceClient) GetTaskById(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/facade.FacadeService/GetTaskById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *facadeServiceClient) UpdateTask(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/facade.FacadeService/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *facadeServiceClient) DeleteTask(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/facade.FacadeService/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *facadeServiceClient) ReadAllTasks(ctx context.Context, in *ReadAllRequest, opts ...grpc.CallOption) (*ReadAllResponse, error) {
	out := new(ReadAllResponse)
	err := c.cc.Invoke(ctx, "/facade.FacadeService/ReadAllTasks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FacadeServiceServer is the server API for FacadeService service.
// All implementations must embed UnimplementedFacadeServiceServer
// for forward compatibility
type FacadeServiceServer interface {
	Signup(context.Context, *SignupRequset) (*SignupResponse, error)
	AllUsers(context.Context, *GetAllUsersRequest) (*GetAlllUsersResponse, error)
	// Create new task
	CreateTask(context.Context, *CreateRequest) (*CreateResponse, error)
	// Read task
	GetTaskById(context.Context, *ReadRequest) (*ReadResponse, error)
	// Update task
	UpdateTask(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete task
	DeleteTask(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Read all tasks
	ReadAllTasks(context.Context, *ReadAllRequest) (*ReadAllResponse, error)
	mustEmbedUnimplementedFacadeServiceServer()
}

// UnimplementedFacadeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFacadeServiceServer struct {
}

func (UnimplementedFacadeServiceServer) Signup(context.Context, *SignupRequset) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedFacadeServiceServer) AllUsers(context.Context, *GetAllUsersRequest) (*GetAlllUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllUsers not implemented")
}
func (UnimplementedFacadeServiceServer) CreateTask(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTask not implemented")
}
func (UnimplementedFacadeServiceServer) GetTaskById(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskById not implemented")
}
func (UnimplementedFacadeServiceServer) UpdateTask(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTask not implemented")
}
func (UnimplementedFacadeServiceServer) DeleteTask(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTask not implemented")
}
func (UnimplementedFacadeServiceServer) ReadAllTasks(context.Context, *ReadAllRequest) (*ReadAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadAllTasks not implemented")
}
func (UnimplementedFacadeServiceServer) mustEmbedUnimplementedFacadeServiceServer() {}

// UnsafeFacadeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FacadeServiceServer will
// result in compilation errors.
type UnsafeFacadeServiceServer interface {
	mustEmbedUnimplementedFacadeServiceServer()
}

func RegisterFacadeServiceServer(s grpc.ServiceRegistrar, srv FacadeServiceServer) {
	s.RegisterService(&FacadeService_ServiceDesc, srv)
}

func _FacadeService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequset)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacadeServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/facade.FacadeService/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacadeServiceServer).Signup(ctx, req.(*SignupRequset))
	}
	return interceptor(ctx, in, info, handler)
}

func _FacadeService_AllUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacadeServiceServer).AllUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/facade.FacadeService/AllUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacadeServiceServer).AllUsers(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FacadeService_CreateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacadeServiceServer).CreateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/facade.FacadeService/CreateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacadeServiceServer).CreateTask(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FacadeService_GetTaskById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacadeServiceServer).GetTaskById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/facade.FacadeService/GetTaskById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacadeServiceServer).GetTaskById(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FacadeService_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacadeServiceServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/facade.FacadeService/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacadeServiceServer).UpdateTask(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FacadeService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacadeServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/facade.FacadeService/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacadeServiceServer).DeleteTask(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FacadeService_ReadAllTasks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FacadeServiceServer).ReadAllTasks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/facade.FacadeService/ReadAllTasks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FacadeServiceServer).ReadAllTasks(ctx, req.(*ReadAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FacadeService_ServiceDesc is the grpc.ServiceDesc for FacadeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FacadeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "facade.FacadeService",
	HandlerType: (*FacadeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _FacadeService_Signup_Handler,
		},
		{
			MethodName: "AllUsers",
			Handler:    _FacadeService_AllUsers_Handler,
		},
		{
			MethodName: "CreateTask",
			Handler:    _FacadeService_CreateTask_Handler,
		},
		{
			MethodName: "GetTaskById",
			Handler:    _FacadeService_GetTaskById_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _FacadeService_UpdateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _FacadeService_DeleteTask_Handler,
		},
		{
			MethodName: "ReadAllTasks",
			Handler:    _FacadeService_ReadAllTasks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/services/facade/proto/facade-service.proto",
}
