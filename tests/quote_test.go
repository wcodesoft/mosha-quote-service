package tests

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-quote-service/data"
	"testing"
)

func TestQuote(t *testing.T) {

	Convey("With QuoteBuilder", t, func() {
		quoteBuilder := data.NewQuoteBuilder()

		Convey("When building a quote", func() {
			quote := quoteBuilder.Build()

			Convey("The quote should be initialized", func() {
				So(quote, ShouldNotBeNil)
			})
		})

		Convey("When building a quote with a specific ID", func() {
			quote := quoteBuilder.WithId("123").Build()

			Convey("The quote should be initialized with the given ID", func() {
				So(quote.ID, ShouldEqual, "123")
			})
		})

		Convey("When building a quote with a specific author ID", func() {
			quote := quoteBuilder.WithAuthorId("John Doe").Build()

			Convey("The quote should be initialized with the given author", func() {
				So(quote.AuthorID, ShouldEqual, "John Doe")
			})
		})

		Convey("When building a quote with a specific text", func() {
			quote := quoteBuilder.WithText("Hello World").Build()

			Convey("The quote should be initialized with the given text", func() {
				So(quote.Text, ShouldEqual, "Hello World")
			})
		})

		Convey("When building a quote with a specific timestamp", func() {
			quote := quoteBuilder.WithTimestamp(123).Build()

			Convey("The quote should be initialized with the given timestamp", func() {
				So(quote.Timestamp, ShouldEqual, 123)
			})
		})

		Convey("Two quotes built with the same builder should be equal", func() {
			quote1 := quoteBuilder.Build()
			quote2 := quoteBuilder.Build()

			Convey("The quotes should be equal", func() {
				So(quote1, ShouldResemble, quote2)
			})
		})
	})
}
