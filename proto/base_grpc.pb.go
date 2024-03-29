// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.3
// source: proto/base.proto

package proto

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

// MensajeServiceClient is the client API for MensajeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MensajeServiceClient interface {
	Create(ctx context.Context, in *Crearmensaje, opts ...grpc.CallOption) (*Respuestamensaje, error)
	CreateLista(ctx context.Context, in *ConsultarLista, opts ...grpc.CallOption) (*RespuestaLista, error)
	CreateActualiza(ctx context.Context, in *CrearActualizacion, opts ...grpc.CallOption) (*RespuestaActualizacion, error)
}

type mensajeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMensajeServiceClient(cc grpc.ClientConnInterface) MensajeServiceClient {
	return &mensajeServiceClient{cc}
}

func (c *mensajeServiceClient) Create(ctx context.Context, in *Crearmensaje, opts ...grpc.CallOption) (*Respuestamensaje, error) {
	out := new(Respuestamensaje)
	err := c.cc.Invoke(ctx, "/grpc.MensajeService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mensajeServiceClient) CreateLista(ctx context.Context, in *ConsultarLista, opts ...grpc.CallOption) (*RespuestaLista, error) {
	out := new(RespuestaLista)
	err := c.cc.Invoke(ctx, "/grpc.MensajeService/CreateLista", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mensajeServiceClient) CreateActualiza(ctx context.Context, in *CrearActualizacion, opts ...grpc.CallOption) (*RespuestaActualizacion, error) {
	out := new(RespuestaActualizacion)
	err := c.cc.Invoke(ctx, "/grpc.MensajeService/CreateActualiza", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MensajeServiceServer is the server API for MensajeService service.
// All implementations must embed UnimplementedMensajeServiceServer
// for forward compatibility
type MensajeServiceServer interface {
	Create(context.Context, *Crearmensaje) (*Respuestamensaje, error)
	CreateLista(context.Context, *ConsultarLista) (*RespuestaLista, error)
	CreateActualiza(context.Context, *CrearActualizacion) (*RespuestaActualizacion, error)
	mustEmbedUnimplementedMensajeServiceServer()
}

// UnimplementedMensajeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMensajeServiceServer struct {
}

func (UnimplementedMensajeServiceServer) Create(context.Context, *Crearmensaje) (*Respuestamensaje, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedMensajeServiceServer) CreateLista(context.Context, *ConsultarLista) (*RespuestaLista, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLista not implemented")
}
func (UnimplementedMensajeServiceServer) CreateActualiza(context.Context, *CrearActualizacion) (*RespuestaActualizacion, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActualiza not implemented")
}
func (UnimplementedMensajeServiceServer) mustEmbedUnimplementedMensajeServiceServer() {}

// UnsafeMensajeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MensajeServiceServer will
// result in compilation errors.
type UnsafeMensajeServiceServer interface {
	mustEmbedUnimplementedMensajeServiceServer()
}

func RegisterMensajeServiceServer(s grpc.ServiceRegistrar, srv MensajeServiceServer) {
	s.RegisterService(&MensajeService_ServiceDesc, srv)
}

func _MensajeService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Crearmensaje)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MensajeServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MensajeService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MensajeServiceServer).Create(ctx, req.(*Crearmensaje))
	}
	return interceptor(ctx, in, info, handler)
}

func _MensajeService_CreateLista_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConsultarLista)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MensajeServiceServer).CreateLista(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MensajeService/CreateLista",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MensajeServiceServer).CreateLista(ctx, req.(*ConsultarLista))
	}
	return interceptor(ctx, in, info, handler)
}

func _MensajeService_CreateActualiza_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CrearActualizacion)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MensajeServiceServer).CreateActualiza(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.MensajeService/CreateActualiza",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MensajeServiceServer).CreateActualiza(ctx, req.(*CrearActualizacion))
	}
	return interceptor(ctx, in, info, handler)
}

// MensajeService_ServiceDesc is the grpc.ServiceDesc for MensajeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MensajeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.MensajeService",
	HandlerType: (*MensajeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _MensajeService_Create_Handler,
		},
		{
			MethodName: "CreateLista",
			Handler:    _MensajeService_CreateLista_Handler,
		},
		{
			MethodName: "CreateActualiza",
			Handler:    _MensajeService_CreateActualiza_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/base.proto",
}
