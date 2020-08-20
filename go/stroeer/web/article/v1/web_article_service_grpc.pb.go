// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package article

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WebArticleServiceClient is the client API for WebArticleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WebArticleServiceClient interface {
	// Returns an Article page including all (SEO) relevant information to
	// render an Article with the given id as canonical web or AMP
	GetWebArticlePage(ctx context.Context, in *GetWebArticlePageRequest, opts ...grpc.CallOption) (*GetWebArticlePageResponse, error)
}

type webArticleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebArticleServiceClient(cc grpc.ClientConnInterface) WebArticleServiceClient {
	return &webArticleServiceClient{cc}
}

func (c *webArticleServiceClient) GetWebArticlePage(ctx context.Context, in *GetWebArticlePageRequest, opts ...grpc.CallOption) (*GetWebArticlePageResponse, error) {
	out := new(GetWebArticlePageResponse)
	err := c.cc.Invoke(ctx, "/stroeer.web.article.v1.WebArticleService/GetWebArticlePage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebArticleServiceServer is the server API for WebArticleService service.
// All implementations must embed UnimplementedWebArticleServiceServer
// for forward compatibility
type WebArticleServiceServer interface {
	// Returns an Article page including all (SEO) relevant information to
	// render an Article with the given id as canonical web or AMP
	GetWebArticlePage(context.Context, *GetWebArticlePageRequest) (*GetWebArticlePageResponse, error)
	mustEmbedUnimplementedWebArticleServiceServer()
}

// UnimplementedWebArticleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWebArticleServiceServer struct {
}

func (*UnimplementedWebArticleServiceServer) GetWebArticlePage(context.Context, *GetWebArticlePageRequest) (*GetWebArticlePageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWebArticlePage not implemented")
}
func (*UnimplementedWebArticleServiceServer) mustEmbedUnimplementedWebArticleServiceServer() {}

func RegisterWebArticleServiceServer(s *grpc.Server, srv WebArticleServiceServer) {
	s.RegisterService(&_WebArticleService_serviceDesc, srv)
}

func _WebArticleService_GetWebArticlePage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWebArticlePageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebArticleServiceServer).GetWebArticlePage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stroeer.web.article.v1.WebArticleService/GetWebArticlePage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebArticleServiceServer).GetWebArticlePage(ctx, req.(*GetWebArticlePageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WebArticleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stroeer.web.article.v1.WebArticleService",
	HandlerType: (*WebArticleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWebArticlePage",
			Handler:    _WebArticleService_GetWebArticlePage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stroeer/web/article/v1/web_article_service.proto",
}