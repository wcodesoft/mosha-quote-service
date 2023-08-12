package repository

import (
	"context"
	"fmt"
	"github.com/wcodesoft/mosha-quote-service/data"
	mdb "github.com/wcodesoft/mosha-service-common/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDatabase struct {
	connection *mdb.MongoConnection
	coll       *mongo.Collection
}

// NewMongoDatabase creates a new mongo database.
func NewMongoDatabase(connection *mdb.MongoConnection) Database {
	return &mongoDatabase{
		connection: connection,
		coll:       connection.Collection,
	}
}

// AddQuote adds a new quote to the database.
func (m mongoDatabase) AddQuote(quote data.Quote) (string, error) {
	if quote.Text == "" {
		panic("Quote text is empty")
	}

	result, err := m.coll.InsertOne(context.Background(), fromQuote(quote))
	if err != nil {
		return "", err
	}
	newId := result.InsertedID
	return fmt.Sprintf("%v", newId), nil
}

// ListAll returns all quotes from the database.
func (m mongoDatabase) ListAll() []data.Quote {
	cursor, err := m.coll.Find(context.Background(), bson.D{})
	if err != nil {
		panic(err)
	}
	var results []quoteDB
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	quotes := make([]data.Quote, len(results))
	for index, v := range results {
		quotes[index] = toQuote(v)
	}
	return quotes
}

// UpdateQuote updates a quote in the database.
func (m mongoDatabase) UpdateQuote(quote data.Quote) (data.Quote, error) {
	filter := bson.D{{"_id", quote.ID}}
	opts := options.Update().SetHint(bson.D{{"_id", 1}})
	update := bson.D{{"$set", fromQuote(quote)}}
	_, err := m.coll.UpdateOne(context.Background(), filter, update, opts)
	if err != nil {
		return data.Quote{}, err
	}
	return quote, nil
}

// DeleteQuote deletes a quote from the database.
func (m mongoDatabase) DeleteQuote(id string) error {
	filter := bson.D{{"_id", id}}
	opts := options.Delete().SetHint(bson.D{{"_id", 1}})
	result, err := m.coll.DeleteOne(context.Background(), filter, opts)
	if result.DeletedCount == 0 {
		return fmt.Errorf("quote with id %s not found", id)
	}
	if err != nil {
		return err
	}
	return nil
}

// GetQuote returns a quote from the database.
func (m mongoDatabase) GetQuote(id string) (data.Quote, error) {
	filter := bson.D{{"_id", id}}
	opts := options.FindOne().SetHint(bson.D{{"_id", 1}})
	var result quoteDB
	err := m.coll.FindOne(context.Background(), filter, opts).Decode(&result)
	if err != nil {
		return data.Quote{}, err
	}
	return toQuote(result), nil

}

// GetAuthorQuotes returns all quotes from an author.
func (m mongoDatabase) GetAuthorQuotes(authorID string) []data.Quote {
	filter := bson.D{{"authorid", authorID}}
	cursor, err := m.coll.Find(context.Background(), filter)
	if err != nil {
		panic(err)
	}
	var results []quoteDB
	if err = cursor.All(context.Background(), &results); err != nil {
		panic(err)
	}
	quotes := make([]data.Quote, len(results))
	for index, v := range results {
		quotes[index] = toQuote(v)
	}
	return quotes
}
