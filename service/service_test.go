package service

import (
	"testing"

	faker "github.com/brianvoe/gofakeit/v6"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-quote-service/data"
	"github.com/wcodesoft/mosha-quote-service/repository"
)

func TestService(t *testing.T) {
	cr := repository.NewFakeClientRepository()
	authorId := "123"
	quote := data.NewQuoteBuilder().
		WithId(faker.UUID()).
		WithText(faker.Quote()).
		WithAuthorId(authorId).
		Build()

	Convey("When creating a new service", t, func() {
		database := repository.NewInMemoryDatabase()
		repo := repository.New(database, cr)
		service := New(repo)

		Convey("The service should be initialized", func() {
			So(service, ShouldNotBeNil)
		})

		Convey("When adding a quote", func() {
			quoteId, _ := service.CreateQuote(quote)
			Convey("The list of quotes should contain the new quote", func() {
				So(len(service.ListAll()), ShouldEqual, 1)
			})

			Convey("Getting the quote by ID should return the correct quote", func() {
				quote, _ := service.GetQuote(quoteId)
				So(quote.ID, ShouldEqual, quoteId)
				So(quote.Text, ShouldEqual, quote.Text)
				So(quote.AuthorID, ShouldEqual, quote.AuthorID)
			})
		})

		Convey("When deleting a quote", func() {
			quoteId, _ := service.CreateQuote(quote)
			err := service.DeleteQuote(quoteId)
			Convey("The list of quotes should be empty", func() {
				So(len(service.ListAll()), ShouldEqual, 0)
			})

			Convey("Getting the quote by ID should return an error", func() {
				_, getErr := service.GetQuote(quoteId)
				So(getErr, ShouldNotBeNil)
			})

			Convey("Deleting the quote should not return an error", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When updating a quote", func() {
			quoteId, _ := service.CreateQuote(quote)
			newText := faker.Quote()
			quote.Text = newText
			_, err := service.UpdateQuote(quote)
			Convey("The list of quotes should contain the updated quote", func() {
				So(len(service.ListAll()), ShouldEqual, 1)
			})

			Convey("Getting the quote by ID should return the correct quote", func() {
				quote, _ := service.GetQuote(quoteId)
				So(quote.ID, ShouldEqual, quoteId)
				So(quote.Text, ShouldEqual, newText)
				So(quote.AuthorID, ShouldEqual, quote.AuthorID)
			})

			Convey("Updating the quote should not return an error", func() {
				So(err, ShouldBeNil)
			})
		})

		Convey("When getting all quotes", func() {
			_, _ = service.CreateQuote(quote)
			quotes := service.ListAll()
			Convey("The list of quotes should not be empty", func() {
				So(len(quotes), ShouldEqual, 1)
			})
		})

		Convey("When getting all quotes by author", func() {
			_, _ = service.CreateQuote(quote)
			quotes := service.GetAuthorQuotes(authorId)
			Convey("The list of quotes should not be empty", func() {
				So(len(quotes), ShouldEqual, 1)
			})
		})

	})
}
