package service

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/wcodesoft/mosha-quote-service/data"
	"github.com/wcodesoft/mosha-quote-service/logger"
	pb "github.com/wcodesoft/mosha-quote-service/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"net"
	"os"
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

func (s server) CreateQuote(_ context.Context, req *pb.CreateQuoteRequest) (*pb.Quote, error) {
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
func (s server) GetQuote(_ context.Context, req *pb.GetQuoteRequest) (*pb.Quote, error) {
	quote, err := s.service.GetQuote(req.Id)
	if err != nil {
		return nil, err
	}
	return toProtoQuote(quote), nil
}

func (s server) UpdateQuote(_ context.Context, req *pb.UpdateQuoteRequest) (*pb.Quote, error) {
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

func (s server) ListQuotes(_ context.Context, _ *emptypb.Empty) (*pb.ListQuotesResponse, error) {
	quotes := s.service.ListAll()
	var pbQuotes []*pb.Quote
	for _, quote := range quotes {
		pbQuotes = append(pbQuotes, toProtoQuote(quote))
	}
	return &pb.ListQuotesResponse{Quotes: pbQuotes}, nil
}

func (s server) DeleteQuote(_ context.Context, request *pb.DeleteQuoteRequest) (*pb.DeleteQuoteResponse, error) {
	err := s.service.DeleteQuote(request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteQuoteResponse{Success: true}, nil
}

func (s server) GetQuotesByAuthor(_ context.Context, request *pb.GetQuotesByAuthorRequest) (*pb.ListQuotesResponse, error) {
	quotes := s.service.GetAuthorQuotes(request.AuthorId)
	var pbQuotes []*pb.Quote
	for _, quote := range quotes {
		pbQuotes = append(pbQuotes, toProtoQuote(quote))
	}
	return &pb.ListQuotesResponse{Quotes: pbQuotes}, nil
}

func (g GrpcRouter) Start(port string) {
	l := log.New(os.Stderr)
	loggerOpts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}
	// Create a new GrpcRouter.
	log.Infof("Starting %s grpc on %s", g.serviceName, port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(logger.InterceptorLogger(l), loggerOpts...),
			// Add logger interceptor to grpc server.
		),
	)
	pb.RegisterQuoteServiceServer(grpcServer, g.server)
	grpcServer.Serve(lis)
}
