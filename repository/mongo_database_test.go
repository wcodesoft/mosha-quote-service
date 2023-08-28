package repository

import (
	faker "github.com/brianvoe/gofakeit/v6"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/wcodesoft/mosha-quote-service/data"
	mdb "github.com/wcodesoft/mosha-service-common/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"testing"
)

const databaseName = "mosha"

func createMockedQuote(id string, text string, timestamp int64, authorId string) bson.D {
	return bson.D{
		{Key: "_id", Value: id},
		{Key: "text", Value: text},
		{Key: "timestamp", Value: timestamp},
		{Key: "authorid", Value: authorId},
	}
}

func TestMongoDB(t *testing.T) {

	Convey("When using a databaseName instance", t, func() {
		mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
		text := faker.Phrase()
		authorId := faker.UUID()
		timestamp := faker.Date().Unix()
		id := faker.UUID()
		defer mt.Close()

		mt.Run("Test AddQuote", func(mt *mtest.T) {
			conn := mdb.NewMongoConnection(mt.Client, databaseName, "quote")
			db := NewMongoDatabase(conn)
			mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "_id", Value: id}})
			Convey("Test AddQuote correctly", mt, func() {
				quote := data.Quote{ID: id, Text: text, Timestamp: timestamp, AuthorID: authorId}
				id, err := db.AddQuote(quote)
				So(err, ShouldBeNil)
				So(id, ShouldEqual, quote.ID)
			})

			mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}, {Key: "_id", Value: id}})
			Convey("Test AddQuote with error", mt, func() {
				quote := data.Quote{ID: id, Text: text, Timestamp: timestamp, AuthorID: authorId}
				id, err := db.AddQuote(quote)
				So(err, ShouldNotBeNil)
				So(id, ShouldEqual, "")
			})
		})

		mt.Run("Test GetQuote", func(mt *mtest.T) {
			conn := mdb.NewMongoConnection(mt.Client, databaseName, "quote")
			db := NewMongoDatabase(conn)
			mockFind := mtest.CreateCursorResponse(
				1,
				"mosha.quotes",
				mtest.FirstBatch,
				createMockedQuote(id, text, timestamp, authorId),
			)
			getCursors := mtest.CreateCursorResponse(0, "mosha.quotes", mtest.NextBatch)
			mt.AddMockResponses(mockFind, getCursors)

			Convey("Test GetQuote correctly", mt, func() {
				quote, err := db.GetQuote(id)
				So(err, ShouldBeNil)
				So(quote.ID, ShouldEqual, id)
				So(quote.Text, ShouldEqual, text)
				So(quote.Timestamp, ShouldEqual, timestamp)
				So(quote.AuthorID, ShouldEqual, authorId)
			})

			Convey("Test GetQuote with error", mt, func() {
				quote, err := db.GetQuote(id)
				So(err, ShouldNotBeNil)
				So(quote.ID, ShouldEqual, "")
				So(quote.Text, ShouldEqual, "")
				So(quote.Timestamp, ShouldEqual, 0)
				So(quote.AuthorID, ShouldEqual, "")
			})
		})

		mt.Run("Test DeleteQuote", func(mt *mtest.T) {
			conn := mdb.NewMongoConnection(mt.Client, databaseName, "quote")
			db := NewMongoDatabase(conn)
			Convey("Test DeleteQuote correctly", mt, func() {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}})
				err := db.DeleteQuote(id)
				So(err, ShouldBeNil)
			})

			Convey("Test DeleteQuote with error", mt, func() {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 0}})
				err := db.DeleteQuote(id)
				So(err, ShouldNotBeNil)
			})
		})

		mt.Run("Test UpdateQuote", func(mt *mtest.T) {
			conn := mdb.NewMongoConnection(mt.Client, databaseName, "quote")
			db := NewMongoDatabase(conn)
			mt.AddMockResponses(bson.D{
				{Key: "ok", Value: 1},
				{Key: "value", Value: createMockedQuote(id, text, timestamp, authorId)}},
			)

			Convey("Test UpdateQuote correctly", mt, func() {
				quote := data.Quote{ID: id, Text: text, Timestamp: timestamp, AuthorID: authorId}
				quote, err := db.UpdateQuote(quote)
				So(err, ShouldBeNil)
				So(quote.ID, ShouldEqual, id)
				So(quote.Text, ShouldEqual, text)
				So(quote.Timestamp, ShouldEqual, timestamp)
				So(quote.AuthorID, ShouldEqual, authorId)
			})

			Convey("Test UpdateQuote with error", mt, func() {
				quote := data.Quote{ID: id, Text: text, Timestamp: timestamp, AuthorID: authorId}
				quote, err := db.UpdateQuote(quote)
				So(err, ShouldNotBeNil)
				So(quote.ID, ShouldEqual, "")
				So(quote.Text, ShouldEqual, "")
				So(quote.Timestamp, ShouldEqual, 0)
				So(quote.AuthorID, ShouldEqual, "")
			})
		})

		mt.Run("Test ListQuotes", func(mt *mtest.T) {
			conn := mdb.NewMongoConnection(mt.Client, databaseName, "quote")
			db := NewMongoDatabase(conn)

			Convey("Test ListAllQuotes correctly", mt, func() {
				otherText := faker.Phrase()
				otherTimestamp := faker.Date().Unix()
				otherId := faker.UUID()

				first := mtest.CreateCursorResponse(
					1,
					"mosha.quotes",
					mtest.FirstBatch,
					createMockedQuote(id, text, timestamp, authorId),
				)

				second := mtest.CreateCursorResponse(
					1,
					"mosha.quotes",
					mtest.NextBatch,
					createMockedQuote(otherId, otherText, otherTimestamp, authorId),
				)

				getCursors := mtest.CreateCursorResponse(0, "mosha.quotes", mtest.NextBatch)
				mt.AddMockResponses(first, second, getCursors)
				quotes := db.ListAll()
				So(len(quotes), ShouldEqual, 2)
			})

			Convey("Test ListAllQuotes with error", mt, func() {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
				quotes := db.ListAll()
				So(len(quotes), ShouldEqual, 0)
			})
		})

		mt.Run("Test GetAuthorQuotes", func(mt *mtest.T) {
			conn := mdb.NewMongoConnection(mt.Client, databaseName, "quote")
			db := NewMongoDatabase(conn)

			Convey("Test GetAuthorQuotes correctly", mt, func() {
				otherText := faker.Phrase()
				otherTimestamp := faker.Date().Unix()
				otherId := faker.UUID()

				first := mtest.CreateCursorResponse(
					1,
					"mosha.quotes",
					mtest.FirstBatch,
					createMockedQuote(id, text, timestamp, authorId),
				)

				second := mtest.CreateCursorResponse(
					1,
					"mosha.quotes",
					mtest.NextBatch,
					createMockedQuote(otherId, otherText, otherTimestamp, authorId),
				)

				getCursors := mtest.CreateCursorResponse(0, "mosha.quotes", mtest.NextBatch)
				mt.AddMockResponses(first, second, getCursors)
				quotes := db.GetAuthorQuotes(authorId)
				So(len(quotes), ShouldEqual, 2)
			})

			Convey("Test GetAuthorQuotes with error", mt, func() {
				mt.AddMockResponses(bson.D{{Key: "ok", Value: 0}})
				quotes := db.GetAuthorQuotes(authorId)
				So(len(quotes), ShouldEqual, 0)
			})
		})

		mt.Run("Test DeleteAuthorQuotes", func(mt *mtest.T) {
			conn := mdb.NewMongoConnection(mt.Client, databaseName, "quote")
			db := NewMongoDatabase(conn)

			Convey("Test DeleteAuthorQuotes correctly", mt, func() {
				otherText := faker.Phrase()
				otherTimestamp := faker.Date().Unix()
				otherId := faker.UUID()

				first := mtest.CreateCursorResponse(
					1,
					"mosha.quotes",
					mtest.FirstBatch,
					createMockedQuote(id, text, timestamp, authorId),
				)

				second := mtest.CreateCursorResponse(
					1,
					"mosha.quotes",
					mtest.NextBatch,
					createMockedQuote(otherId, otherText, otherTimestamp, authorId),
				)

				getCursors := mtest.CreateCursorResponse(0, "mosha.quotes", mtest.NextBatch)
				deleteResponse := mtest.CreateSuccessResponse()
				mt.AddMockResponses(first, second, getCursors, deleteResponse)
				quotes := db.GetAuthorQuotes(authorId)
				So(len(quotes), ShouldEqual, 2)
				err := db.DeleteAuthorQuotes(authorId)
				So(err, ShouldBeNil)
			})
		})

		mt.Run("Test GetRandomQuote", func(mt *mtest.T) {
			conn := mdb.NewMongoConnection(mt.Client, databaseName, "quote")
			db := NewMongoDatabase(conn)
			mockFind := mtest.CreateCursorResponse(
				1,
				"mosha.quotes",
				mtest.FirstBatch,
				createMockedQuote(id, text, timestamp, authorId),
			)
			getCursors := mtest.CreateCursorResponse(0, "mosha.quotes", mtest.NextBatch)
			mt.AddMockResponses(mockFind, getCursors)
			Convey("Test GetRandomQuote correctly", mt, func() {
				quote, err := db.GetRandomQuote()
				So(err, ShouldBeNil)
				So(quote.ID, ShouldEqual, id)
			})
		})
	})
}
