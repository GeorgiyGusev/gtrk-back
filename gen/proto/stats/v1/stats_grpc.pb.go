// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: proto/stats/v1/stats.proto

package stats_gen_v1

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
	NewsStatisticsService_GetViewsStatisticsForNews_FullMethodName    = "/proto.stats.v1.NewsStatisticsService/GetViewsStatisticsForNews"
	NewsStatisticsService_GetViewsStatisticsForAllNews_FullMethodName = "/proto.stats.v1.NewsStatisticsService/GetViewsStatisticsForAllNews"
)

// NewsStatisticsServiceClient is the client API for NewsStatisticsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис для получения статистики по новостям
type NewsStatisticsServiceClient interface {
	// Метод для получения статистики по просмотрам конкретной новости
	GetViewsStatisticsForNews(ctx context.Context, in *GetViewsStatisticsForNewsRequest, opts ...grpc.CallOption) (*GetViewsStatisticsForNewsResponse, error)
	// Метод для получения статистики по просмотрам всех новостей
	GetViewsStatisticsForAllNews(ctx context.Context, in *GetViewsStatisticsForAllNewsRequest, opts ...grpc.CallOption) (*GetViewsStatisticsForAllNewsResponse, error)
}

type newsStatisticsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsStatisticsServiceClient(cc grpc.ClientConnInterface) NewsStatisticsServiceClient {
	return &newsStatisticsServiceClient{cc}
}

func (c *newsStatisticsServiceClient) GetViewsStatisticsForNews(ctx context.Context, in *GetViewsStatisticsForNewsRequest, opts ...grpc.CallOption) (*GetViewsStatisticsForNewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetViewsStatisticsForNewsResponse)
	err := c.cc.Invoke(ctx, NewsStatisticsService_GetViewsStatisticsForNews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsStatisticsServiceClient) GetViewsStatisticsForAllNews(ctx context.Context, in *GetViewsStatisticsForAllNewsRequest, opts ...grpc.CallOption) (*GetViewsStatisticsForAllNewsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetViewsStatisticsForAllNewsResponse)
	err := c.cc.Invoke(ctx, NewsStatisticsService_GetViewsStatisticsForAllNews_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsStatisticsServiceServer is the server API for NewsStatisticsService service.
// All implementations must embed UnimplementedNewsStatisticsServiceServer
// for forward compatibility.
//
// Сервис для получения статистики по новостям
type NewsStatisticsServiceServer interface {
	// Метод для получения статистики по просмотрам конкретной новости
	GetViewsStatisticsForNews(context.Context, *GetViewsStatisticsForNewsRequest) (*GetViewsStatisticsForNewsResponse, error)
	// Метод для получения статистики по просмотрам всех новостей
	GetViewsStatisticsForAllNews(context.Context, *GetViewsStatisticsForAllNewsRequest) (*GetViewsStatisticsForAllNewsResponse, error)
	mustEmbedUnimplementedNewsStatisticsServiceServer()
}

// UnimplementedNewsStatisticsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNewsStatisticsServiceServer struct{}

func (UnimplementedNewsStatisticsServiceServer) GetViewsStatisticsForNews(context.Context, *GetViewsStatisticsForNewsRequest) (*GetViewsStatisticsForNewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetViewsStatisticsForNews not implemented")
}
func (UnimplementedNewsStatisticsServiceServer) GetViewsStatisticsForAllNews(context.Context, *GetViewsStatisticsForAllNewsRequest) (*GetViewsStatisticsForAllNewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetViewsStatisticsForAllNews not implemented")
}
func (UnimplementedNewsStatisticsServiceServer) mustEmbedUnimplementedNewsStatisticsServiceServer() {}
func (UnimplementedNewsStatisticsServiceServer) testEmbeddedByValue()                               {}

// UnsafeNewsStatisticsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsStatisticsServiceServer will
// result in compilation errors.
type UnsafeNewsStatisticsServiceServer interface {
	mustEmbedUnimplementedNewsStatisticsServiceServer()
}

func RegisterNewsStatisticsServiceServer(s grpc.ServiceRegistrar, srv NewsStatisticsServiceServer) {
	// If the following call pancis, it indicates UnimplementedNewsStatisticsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NewsStatisticsService_ServiceDesc, srv)
}

func _NewsStatisticsService_GetViewsStatisticsForNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetViewsStatisticsForNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsStatisticsServiceServer).GetViewsStatisticsForNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsStatisticsService_GetViewsStatisticsForNews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsStatisticsServiceServer).GetViewsStatisticsForNews(ctx, req.(*GetViewsStatisticsForNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NewsStatisticsService_GetViewsStatisticsForAllNews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetViewsStatisticsForAllNewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsStatisticsServiceServer).GetViewsStatisticsForAllNews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NewsStatisticsService_GetViewsStatisticsForAllNews_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsStatisticsServiceServer).GetViewsStatisticsForAllNews(ctx, req.(*GetViewsStatisticsForAllNewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NewsStatisticsService_ServiceDesc is the grpc.ServiceDesc for NewsStatisticsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NewsStatisticsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.stats.v1.NewsStatisticsService",
	HandlerType: (*NewsStatisticsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetViewsStatisticsForNews",
			Handler:    _NewsStatisticsService_GetViewsStatisticsForNews_Handler,
		},
		{
			MethodName: "GetViewsStatisticsForAllNews",
			Handler:    _NewsStatisticsService_GetViewsStatisticsForAllNews_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/stats/v1/stats.proto",
}
