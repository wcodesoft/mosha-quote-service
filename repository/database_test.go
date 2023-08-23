package repository

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-quote-service/data"
	"testing"
	"time"
)

func TestDatabase(t *testing.T) {
	Convey("When converting quote http to databaseName model", t, func() {
		quote := quoteDB{ID: "ID", AuthorID: "AuthorID", Text: "Text", Timestamp: time.Now().UnixMilli()}
		quoteHttp := toQuote(quote)
		So(quoteHttp.ID, ShouldEqual, quote.ID)
		So(quoteHttp.AuthorID, ShouldEqual, quote.AuthorID)
		So(quoteHttp.Text, ShouldEqual, quote.Text)
		So(quoteHttp.Timestamp, ShouldEqual, quote.Timestamp)
	})

	Convey("When converting quote databaseName to http model", t, func() {
		quote := data.Quote{ID: "ID", AuthorID: "AuthorID", Text: "Text", Timestamp: time.Now().UnixMilli()}
		quoteDb := fromQuote(quote)
		So(quoteDb.ID, ShouldEqual, quote.ID)
		So(quoteDb.AuthorID, ShouldEqual, quote.AuthorID)
		So(quoteDb.Text, ShouldEqual, quote.Text)
		So(quoteDb.Timestamp, ShouldEqual, quote.Timestamp)
	})
}
