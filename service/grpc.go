package service

import (
	"context"
	"fmt"
	"github.com/wcodesoft/mosha-quote-service/data"
	pb "github.com/wcodesoft/mosha-quote-service/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// GrpcRouter represents the gRPC router.
type GrpcRouter struct {
	service Service
	pb.UnimplementedQuoteServiceServer
}

func toProtoQuote(quote data.Quote) *pb.Quote {
	return &pb.Quote{Id: quote.ID, AuthorId: quote.AuthorID, Text: quote.Text, Timestamp: quote.Timestamp}
}

func toQuoteDB(author *pb.Quote) data.Quote {
	return data.Quote{ID: author.Id, AuthorID: author.AuthorId, Text: author.Text, Timestamp: author.Timestamp}
}

// NewGrpcRouter creates a new gRPC router.
func NewGrpcRouter(s Service) pb.QuoteServiceServer {
	return GrpcRouter{service: s}
}

func (s GrpcRouter) CreateQuote(_ context.Context, req *pb.CreateQuoteRequest) (*pb.Quote, error) {
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
func (s GrpcRouter) GetQuote(_ context.Context, req *pb.GetQuoteRequest) (*pb.Quote, error) {
	quote, err := s.service.GetQuote(req.Id)
	if err != nil {
		return nil, err
	}
	return toProtoQuote(quote), nil
}

func (s GrpcRouter) UpdateQuote(_ context.Context, req *pb.UpdateQuoteRequest) (*pb.Quote, error) {
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

func (s GrpcRouter) ListQuotes(_ context.Context, _ *emptypb.Empty) (*pb.ListQuotesResponse, error) {
	quotes := s.service.ListAll()
	var pbQuotes []*pb.Quote
	for _, quote := range quotes {
		pbQuotes = append(pbQuotes, toProtoQuote(quote))
	}
	return &pb.ListQuotesResponse{Quotes: pbQuotes}, nil
}

func (s GrpcRouter) DeleteQuote(_ context.Context, request *pb.DeleteQuoteRequest) (*pb.DeleteQuoteResponse, error) {
	err := s.service.DeleteQuote(request.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteQuoteResponse{Success: true}, nil
}

func (s GrpcRouter) GetQuotesByAuthor(_ context.Context, request *pb.GetQuotesByAuthorRequest) (*pb.ListQuotesResponse, error) {
	quotes := s.service.GetAuthorQuotes(request.AuthorId)
	var pbQuotes []*pb.Quote
	for _, quote := range quotes {
		pbQuotes = append(pbQuotes, toProtoQuote(quote))
	}
	return &pb.ListQuotesResponse{Quotes: pbQuotes}, nil
}
