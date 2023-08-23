package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	faker "github.com/brianvoe/gofakeit/v6"
	. "github.com/smartystreets/goconvey/convey"
	pb "github.com/wcodesoft/mosha-quote-service/proto"
	"github.com/wcodesoft/mosha-quote-service/repository"
	"testing"
	"time"
)

func createGrpcRouter() GrpcRouter {
	memoryDatabase := repository.NewInMemoryDatabase()
	cr := repository.NewFakeClientRepository()
	repo := repository.New(memoryDatabase, cr)
	service := New(repo)
	router := NewGrpcRouter(service, "AuthorService")
	return router
}

func TestGrpc(t *testing.T) {
	id := faker.UUID()
	text := faker.Quote()
	newText := faker.Quote()
	authorId := "123"
	timestamp := time.Now().UnixMilli()
	quote := &pb.Quote{
		Id:        id,
		Text:      text,
		AuthorId:  authorId,
		Timestamp: timestamp,
	}

	updateQuote := &pb.Quote{
		Id:        id,
		Text:      newText,
		AuthorId:  authorId,
		Timestamp: timestamp,
	}

	Convey("When adding valid quote", t, func() {
		router := createGrpcRouter()
		res, err := router.server.CreateQuote(context.Background(),
			&pb.CreateQuoteRequest{Quote: quote},
		)
		Convey("The response should not be nil", func() {
			So(res, ShouldNotBeNil)
		})
		Convey("The error should be nil", func() {
			So(err, ShouldBeNil)
		})
		Convey("The response should contain the correct ID", func() {
			So(res.Id, ShouldEqual, id)
		})
	})

	Convey("With a quote in the database", t, func() {
		router := createGrpcRouter()
		router.server.CreateQuote(context.Background(),
			&pb.CreateQuoteRequest{Quote: quote},
		)

		Convey("When getting the quote", func() {
			res, err := router.server.GetQuote(context.Background(),
				&pb.GetQuoteRequest{Id: id},
			)
			Convey("The response should not be nil", func() {
				So(res, ShouldNotBeNil)
			})
			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The response should contain the correct ID", func() {
				So(res.Id, ShouldEqual, id)
			})
		})

		Convey("When deleting the quote", func() {
			res, err := router.server.DeleteQuote(context.Background(),
				&pb.DeleteQuoteRequest{Id: id},
			)
			Convey("The response should not be nil", func() {
				So(res, ShouldNotBeNil)
			})
			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The response should contain the correct ID", func() {
				So(res.Success, ShouldEqual, true)
			})
		})

		Convey("When updating the quote", func() {
			res, err := router.server.UpdateQuote(context.Background(),
				&pb.UpdateQuoteRequest{Quote: updateQuote},
			)
			Convey("The response should not be nil", func() {
				So(res, ShouldNotBeNil)
			})
			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The response should contain the correct ID", func() {
				So(res.Id, ShouldEqual, id)
			})
			Convey("The response should contain the correct text", func() {
				So(res.Text, ShouldEqual, newText)
			})
		})

		Convey("When getting all quotes", func() {
			res, err := router.server.ListQuotes(context.Background(), &emptypb.Empty{})
			Convey("The response should not be nil", func() {
				So(res, ShouldNotBeNil)
			})
			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The response should contain the correct quantity of quotes", func() {
				So(len(res.Quotes), ShouldEqual, 1)
			})
		})

		Convey("When getting all quotes by author", func() {
			res, err := router.server.GetQuotesByAuthor(context.Background(),
				&pb.GetQuotesByAuthorRequest{AuthorId: authorId},
			)
			Convey("The response should not be nil", func() {
				So(res, ShouldNotBeNil)
			})
			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The response should contain the correct quantity of quotes", func() {
				So(len(res.Quotes), ShouldEqual, 1)
			})
		})

		Convey("When getting quote by author that does not exist", func() {
			res, err := router.server.GetQuotesByAuthor(context.Background(),
				&pb.GetQuotesByAuthorRequest{AuthorId: "1234"},
			)
			Convey("The response should not be nil", func() {
				So(res, ShouldNotBeNil)
			})
			Convey("The error should be nil", func() {
				So(err, ShouldBeNil)
			})
			Convey("The response should contain the correct quantity of quotes", func() {
				So(len(res.Quotes), ShouldEqual, 0)
			})
		})
	})

	Convey("When database is empty", t, func() {
		router := createGrpcRouter()

		Convey("When updating a quote that does not exist", func() {
			res, err := router.server.UpdateQuote(context.Background(),
				&pb.UpdateQuoteRequest{Quote: updateQuote},
			)
			Convey("The response should be nil", func() {
				So(res, ShouldBeNil)
			})
			Convey("The error should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When getting a quote that does not exist", func() {
			res, err := router.server.GetQuote(context.Background(),
				&pb.GetQuoteRequest{Id: id},
			)
			Convey("The response should be nil", func() {
				So(res, ShouldBeNil)
			})
			Convey("The error should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When deleting a quote that does not exist", func() {
			res, err := router.server.DeleteQuote(context.Background(),
				&pb.DeleteQuoteRequest{Id: id},
			)
			Convey("The response should be nil", func() {
				So(res, ShouldBeNil)
			})
			Convey("The error should not be nil", func() {
				So(err, ShouldNotBeNil)
			})
		})
	})

	Convey("When adding a quote with an invalid author", t, func() {
		router := createGrpcRouter()
		quote.AuthorId = "1234"
		res, err := router.server.CreateQuote(context.Background(),
			&pb.CreateQuoteRequest{Quote: quote},
		)
		Convey("The response should be nil", func() {
			So(res, ShouldBeNil)
		})
		Convey("The error should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})

	Convey("When adding invalid quote", t, func() {
		router := createGrpcRouter()
		res, err := router.server.CreateQuote(context.Background(),
			&pb.CreateQuoteRequest{Quote: nil},
		)
		Convey("The response should be nil", func() {
			So(res, ShouldBeNil)
		})
		Convey("The error should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})

	Convey("When updating a quote with an invalid author", t, func() {
		router := createGrpcRouter()
		res, err := router.server.UpdateQuote(context.Background(),
			&pb.UpdateQuoteRequest{Quote: nil},
		)
		Convey("The response should be nil", func() {
			So(res, ShouldBeNil)
		})
		Convey("The error should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})
}
