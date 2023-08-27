// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: proto/quote.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	QuoteService_GetQuote_FullMethodName                = "/quoteservice.QuoteService/GetQuote"
	QuoteService_ListQuotes_FullMethodName              = "/quoteservice.QuoteService/ListQuotes"
	QuoteService_CreateQuote_FullMethodName             = "/quoteservice.QuoteService/CreateQuote"
	QuoteService_UpdateQuote_FullMethodName             = "/quoteservice.QuoteService/UpdateQuote"
	QuoteService_DeleteQuote_FullMethodName             = "/quoteservice.QuoteService/DeleteQuote"
	QuoteService_GetQuotesByAuthor_FullMethodName       = "/quoteservice.QuoteService/GetQuotesByAuthor"
	QuoteService_DeleteAllQuotesByAuthor_FullMethodName = "/quoteservice.QuoteService/DeleteAllQuotesByAuthor"
)

// QuoteServiceClient is the client API for QuoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuoteServiceClient interface {
	// GetQuote returns a quote by id
	GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...grpc.CallOption) (*Quote, error)
	// ListQuotes returns a list of quotes
	ListQuotes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListQuotesResponse, error)
	// CreateQuote creates a new quote
	CreateQuote(ctx context.Context, in *CreateQuoteRequest, opts ...grpc.CallOption) (*Quote, error)
	// UpdateQuote updates a quote
	UpdateQuote(ctx context.Context, in *UpdateQuoteRequest, opts ...grpc.CallOption) (*Quote, error)
	// DeleteQuote deletes a quote
	DeleteQuote(ctx context.Context, in *DeleteQuoteRequest, opts ...grpc.CallOption) (*DeleteQuoteResponse, error)
	// GetQuoteByAuthor returns a list of quotes by author id
	GetQuotesByAuthor(ctx context.Context, in *GetQuotesByAuthorRequest, opts ...grpc.CallOption) (*ListQuotesResponse, error)
	// DeleteAllQuotesByAuthor deletes all quotes by author id
	DeleteAllQuotesByAuthor(ctx context.Context, in *DeleteQuotesByAuthorRequest, opts ...grpc.CallOption) (*DeleteQuoteResponse, error)
}

type quoteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuoteServiceClient(cc grpc.ClientConnInterface) QuoteServiceClient {
	return &quoteServiceClient{cc}
}

func (c *quoteServiceClient) GetQuote(ctx context.Context, in *GetQuoteRequest, opts ...grpc.CallOption) (*Quote, error) {
	out := new(Quote)
	err := c.cc.Invoke(ctx, QuoteService_GetQuote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) ListQuotes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListQuotesResponse, error) {
	out := new(ListQuotesResponse)
	err := c.cc.Invoke(ctx, QuoteService_ListQuotes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) CreateQuote(ctx context.Context, in *CreateQuoteRequest, opts ...grpc.CallOption) (*Quote, error) {
	out := new(Quote)
	err := c.cc.Invoke(ctx, QuoteService_CreateQuote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) UpdateQuote(ctx context.Context, in *UpdateQuoteRequest, opts ...grpc.CallOption) (*Quote, error) {
	out := new(Quote)
	err := c.cc.Invoke(ctx, QuoteService_UpdateQuote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) DeleteQuote(ctx context.Context, in *DeleteQuoteRequest, opts ...grpc.CallOption) (*DeleteQuoteResponse, error) {
	out := new(DeleteQuoteResponse)
	err := c.cc.Invoke(ctx, QuoteService_DeleteQuote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) GetQuotesByAuthor(ctx context.Context, in *GetQuotesByAuthorRequest, opts ...grpc.CallOption) (*ListQuotesResponse, error) {
	out := new(ListQuotesResponse)
	err := c.cc.Invoke(ctx, QuoteService_GetQuotesByAuthor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *quoteServiceClient) DeleteAllQuotesByAuthor(ctx context.Context, in *DeleteQuotesByAuthorRequest, opts ...grpc.CallOption) (*DeleteQuoteResponse, error) {
	out := new(DeleteQuoteResponse)
	err := c.cc.Invoke(ctx, QuoteService_DeleteAllQuotesByAuthor_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuoteServiceServer is the server API for QuoteService service.
// All implementations must embed UnimplementedQuoteServiceServer
// for forward compatibility
type QuoteServiceServer interface {
	// GetQuote returns a quote by id
	GetQuote(context.Context, *GetQuoteRequest) (*Quote, error)
	// ListQuotes returns a list of quotes
	ListQuotes(context.Context, *emptypb.Empty) (*ListQuotesResponse, error)
	// CreateQuote creates a new quote
	CreateQuote(context.Context, *CreateQuoteRequest) (*Quote, error)
	// UpdateQuote updates a quote
	UpdateQuote(context.Context, *UpdateQuoteRequest) (*Quote, error)
	// DeleteQuote deletes a quote
	DeleteQuote(context.Context, *DeleteQuoteRequest) (*DeleteQuoteResponse, error)
	// GetQuoteByAuthor returns a list of quotes by author id
	GetQuotesByAuthor(context.Context, *GetQuotesByAuthorRequest) (*ListQuotesResponse, error)
	// DeleteAllQuotesByAuthor deletes all quotes by author id
	DeleteAllQuotesByAuthor(context.Context, *DeleteQuotesByAuthorRequest) (*DeleteQuoteResponse, error)
	mustEmbedUnimplementedQuoteServiceServer()
}

// UnimplementedQuoteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQuoteServiceServer struct {
}

func (UnimplementedQuoteServiceServer) GetQuote(context.Context, *GetQuoteRequest) (*Quote, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuote not implemented")
}
func (UnimplementedQuoteServiceServer) ListQuotes(context.Context, *emptypb.Empty) (*ListQuotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQuotes not implemented")
}
func (UnimplementedQuoteServiceServer) CreateQuote(context.Context, *CreateQuoteRequest) (*Quote, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQuote not implemented")
}
func (UnimplementedQuoteServiceServer) UpdateQuote(context.Context, *UpdateQuoteRequest) (*Quote, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuote not implemented")
}
func (UnimplementedQuoteServiceServer) DeleteQuote(context.Context, *DeleteQuoteRequest) (*DeleteQuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQuote not implemented")
}
func (UnimplementedQuoteServiceServer) GetQuotesByAuthor(context.Context, *GetQuotesByAuthorRequest) (*ListQuotesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQuotesByAuthor not implemented")
}
func (UnimplementedQuoteServiceServer) DeleteAllQuotesByAuthor(context.Context, *DeleteQuotesByAuthorRequest) (*DeleteQuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllQuotesByAuthor not implemented")
}
func (UnimplementedQuoteServiceServer) mustEmbedUnimplementedQuoteServiceServer() {}

// UnsafeQuoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuoteServiceServer will
// result in compilation errors.
type UnsafeQuoteServiceServer interface {
	mustEmbedUnimplementedQuoteServiceServer()
}

func RegisterQuoteServiceServer(s grpc.ServiceRegistrar, srv QuoteServiceServer) {
	s.RegisterService(&QuoteService_ServiceDesc, srv)
}

func _QuoteService_GetQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).GetQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuoteService_GetQuote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).GetQuote(ctx, req.(*GetQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_ListQuotes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).ListQuotes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuoteService_ListQuotes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).ListQuotes(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_CreateQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).CreateQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuoteService_CreateQuote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).CreateQuote(ctx, req.(*CreateQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_UpdateQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).UpdateQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuoteService_UpdateQuote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).UpdateQuote(ctx, req.(*UpdateQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_DeleteQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).DeleteQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuoteService_DeleteQuote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).DeleteQuote(ctx, req.(*DeleteQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_GetQuotesByAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQuotesByAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).GetQuotesByAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuoteService_GetQuotesByAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).GetQuotesByAuthor(ctx, req.(*GetQuotesByAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuoteService_DeleteAllQuotesByAuthor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQuotesByAuthorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuoteServiceServer).DeleteAllQuotesByAuthor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuoteService_DeleteAllQuotesByAuthor_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuoteServiceServer).DeleteAllQuotesByAuthor(ctx, req.(*DeleteQuotesByAuthorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuoteService_ServiceDesc is the grpc.ServiceDesc for QuoteService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuoteService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "quoteservice.QuoteService",
	HandlerType: (*QuoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetQuote",
			Handler:    _QuoteService_GetQuote_Handler,
		},
		{
			MethodName: "ListQuotes",
			Handler:    _QuoteService_ListQuotes_Handler,
		},
		{
			MethodName: "CreateQuote",
			Handler:    _QuoteService_CreateQuote_Handler,
		},
		{
			MethodName: "UpdateQuote",
			Handler:    _QuoteService_UpdateQuote_Handler,
		},
		{
			MethodName: "DeleteQuote",
			Handler:    _QuoteService_DeleteQuote_Handler,
		},
		{
			MethodName: "GetQuotesByAuthor",
			Handler:    _QuoteService_GetQuotesByAuthor_Handler,
		},
		{
			MethodName: "DeleteAllQuotesByAuthor",
			Handler:    _QuoteService_DeleteAllQuotesByAuthor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/quote.proto",
}
