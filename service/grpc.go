package service

import (
	"context"
	"fmt"
	"github.com/wcodesoft/mosha-quote-service/data"
	pb "github.com/wcodesoft/mosha-quote-service/proto"
	"github.com/wcodesoft/mosha-service-common/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
)

// GrpcRouter represents the gRPC router.
type GrpcRouter struct {
	serviceName string
	server      pb.QuoteServiceServer
}

type server struct {
	service Service
	pb.UnimplementedQuoteServiceServer
}

func newServer(s Service) pb.QuoteServiceServer {
	return &server{
		service: s,
	}
}

func toProtoQuote(quote data.Quote) *pb.Quote {
	return &pb.Quote{Id: quote.ID, AuthorId: quote.AuthorID, Text: quote.Text, Timestamp: quote.Timestamp}
}

func toQuoteDB(author *pb.Quote) data.Quote {
	return data.Quote{ID: author.Id, AuthorID: author.AuthorId, Text: author.Text, Timestamp: author.Timestamp}
}

// NewGrpcRouter creates a new gRPC router.
func NewGrpcRouter(s Service, serviceName string) GrpcRouter {
	return GrpcRouter{
		server:      newServer(s),
		serviceName: serviceName}
}

func (s *server) CreateQuote(_ context.Context, req *pb.CreateQuoteRequest) (*pb.Quote, error) {
	quote := req.GetQuote()
	if quote == nil {
		return nil, fmt.Errorf("quote is nil")
	}
	id, err := s.service.CreateQuote(toQuoteDB(quote))
	if err != nil {
		return nil, err
	}
	return &pb.Quote{
		Id:        id,
		AuthorId:  quote.AuthorId,
		Text:      quote.Text,
		Timestamp: quote.Timestamp,
	}, nil
}

func (s *server) GetQuote(_ context.Context, req *pb.GetQuoteRequest) (*pb.Quote, error) {
	quote, err := s.service.GetQuote(req.Id)
	if err != nil {
		return nil, err
	}
	return toProtoQuote(quote), nil
}

func (s *server) UpdateQuote(_ context.Context, req *pb.UpdateQuoteRequest) (*pb.Quote, error) {
	quote := req.GetQuote()
	if quote == nil {
		return nil, fmt.Errorf("quote is nil")
	}
	updatedQuote, err := s.service.UpdateQuote(toQuoteDB(quote))
	if err != nil {
		return nil, err
	}
	return toProtoQuote(updatedQuote), nil
}

func (s *server) ListQuotes(_ context.Context, _ *emptypb.Empty) (*pb.ListQuotesResponse, error) {
	quotes := s.service.ListAll()
	var pbQuotes []*pb.Quote
	for _, quote := range quotes {
		pbQuotes = append(pbQuotes, toProtoQuote(quote))
	}
	return &pb.ListQuotesResponse{Quotes: pbQuotes}, nil
}

func (s *server) DeleteQuote(_ context.Context, request *pb.DeleteQuoteRequest) (*pb.DeleteQuoteResponse, error) {
	err := s.service.DeleteQuote(request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteQuoteResponse{Success: true}, nil
}

func (s *server) GetQuotesByAuthor(_ context.Context, request *pb.GetQuotesByAuthorRequest) (*pb.ListQuotesResponse, error) {
	quotes := s.service.GetAuthorQuotes(request.AuthorId)
	var pbQuotes []*pb.Quote
	for _, quote := range quotes {
		pbQuotes = append(pbQuotes, toProtoQuote(quote))
	}
	return &pb.ListQuotesResponse{Quotes: pbQuotes}, nil
}

func (s *server) DeleteAllQuotesByAuthor(_ context.Context, request *pb.DeleteQuotesByAuthorRequest) (*pb.DeleteQuoteResponse, error) {
	err := s.service.DeleteAuthorQuotes(request.AuthorId)
	return &pb.DeleteQuoteResponse{Success: err != nil}, err
}

func (g GrpcRouter) Start(port string) error {
	grpcServer := grpc.CreateNewGRPCServer()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}
	pb.RegisterQuoteServiceServer(grpcServer, g.server)
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %v", err)
	}
	return nil
}
