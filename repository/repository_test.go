package repository

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-quote-service/data"
	"testing"
)

func TestRepository(t *testing.T) {

	Convey("With repository", t, func() {
		db := NewInMemoryDatabase()
		repo := New(db, NewFakeClientRepository())

		Convey("When adding a quote", func() {
			quote := data.NewQuoteBuilder().WithId("123").WithAuthorId("123").Build()
			id, _ := repo.AddQuote(quote)

			Convey("The list of quotes should contain the new quote", func() {
				quotes := repo.ListAll()
				So(quotes, ShouldContain, quote)
			})

			Convey("Getting the quote by ID should return the correct quote", func() {
				quote, _ := repo.GetQuote(id)
				So(quote.ID, ShouldEqual, "123")
				So(quote.AuthorID, ShouldEqual, "123")
			})

			Convey("Updating the quote should return the updated quote", func() {
				quote, _ := repo.UpdateQuote(
					data.NewQuoteBuilder().
						WithId(id).
						WithAuthorId("456").
						Build(),
				)
				So(quote.ID, ShouldEqual, "123")
				So(quote.AuthorID, ShouldNotEqual, "123")
				So(quote.AuthorID, ShouldEqual, "456")
			})

			Convey("Adding with same ID should fail", func() {
				_, err := repo.AddQuote(quote)
				So(err, ShouldNotBeNil)
			})

			Convey("Getting all quotes by author ID should return the correct quotes", func() {
				quotes := repo.GetAuthorQuotes("123")
				So(quotes, ShouldContain, quote)
				So(len(quotes), ShouldEqual, 1)
			})

			Convey("Deleting all quotes by author ID should delete the correct quotes", func() {
				err := repo.DeleteAuthorQuotes("123")
				So(err, ShouldBeNil)
				quotes := repo.GetAuthorQuotes("123")
				So(len(quotes), ShouldEqual, 0)
			})

			Convey("Getting a random quote should return a quote", func() {
				quote, err := repo.GetRandomQuote()
				So(err, ShouldBeNil)
				So(quote, ShouldNotBeNil)
			})
		})

		Convey("When adding a quote with an invalid author ID", func() {

			Convey("Adding the quote should fail", func() {
				_, err := repo.AddQuote(data.NewQuoteBuilder().WithId("123").WithAuthorId("000").Build())
				So(err, ShouldNotBeNil)
			})
		})

		Convey("When deleting a quote", func() {
			quote := data.NewQuoteBuilder().WithId("123").WithAuthorId("123").Build()
			id, _ := repo.AddQuote(quote)

			Convey("Deleting the quote should remove it from the list", func() {
				if err := repo.DeleteQuote(id); err != nil {
					t.Fatal(err)
				}
				So(repo.ListAll(), ShouldNotContain, quote)
			})
		})

		Convey("When updating a quote with an invalid author ID", func() {
			Convey("Updating the quote should fail", func() {
				_, err := repo.UpdateQuote(data.NewQuoteBuilder().WithId("123").WithAuthorId("000").Build())
				So(err, ShouldNotBeNil)
			})
		})

	})
}
