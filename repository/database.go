package repository

import "github.com/wcodesoft/mosha-quote-service/data"

// Database represents the databaseName interface.
type Database interface {
	// AddQuote adds a new quote to the database.
	AddQuote(quote data.Quote) (string, error)

	// ListAll returns all quotes in the database.
	ListAll() []data.Quote

	// UpdateQuote updates a quote in the database.
	UpdateQuote(quote data.Quote) (data.Quote, error)

	// DeleteQuote deletes a quote from the database.
	DeleteQuote(id string) error

	// GetQuote returns a quote from the database.
	GetQuote(id string) (data.Quote, error)

	// GetAuthorQuotes returns all quotes from an author.
	GetAuthorQuotes(authorID string) []data.Quote

	// DeleteAuthorQuotes deletes all quotes from an author.
	DeleteAuthorQuotes(authorID string) error

	// GetRandomQuote returns a random quote from the database.
	GetRandomQuote() (data.Quote, error)
}

type quoteDB struct {
	ID        string `bson:"_id" json:"id,omitempty"`
	AuthorID  string `bson:"authorid"`
	Text      string `bson:"text"`
	Timestamp int64  `bson:"timestamp"`
}

func fromQuote(quote data.Quote) quoteDB {
	return quoteDB{
		ID:        quote.ID,
		AuthorID:  quote.AuthorID,
		Text:      quote.Text,
		Timestamp: quote.Timestamp,
	}
}

func toQuote(qdb quoteDB) data.Quote {
	return data.Quote{
		ID:        qdb.ID,
		AuthorID:  qdb.AuthorID,
		Text:      qdb.Text,
		Timestamp: qdb.Timestamp,
	}
}
