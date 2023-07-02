package data

import (
	"github.com/google/uuid"
	"time"
)

// Quote represents a quote.
type Quote struct {
	// ID is the unique identifier of the quote.
	ID string `json:"id"`
	//	AuthorID is the unique identifier of the author.
	AuthorID string `json:"authorId"`
	//	Text of the quote.
	Text string `json:"text"`
	// Timestamp is the timestamp when the quote was added.
	Timestamp int64 `json:"timestamp"`
}

// QuoteBuilder is the interface that builds a quote.
type QuoteBuilder interface {
	WithId(id string) QuoteBuilder
	WithAuthorId(authorId string) QuoteBuilder
	WithText(text string) QuoteBuilder
	WithTimestamp(timestamp int64) QuoteBuilder
	Build() Quote
}

type quoteBuilder struct {
	id        string
	authorId  string
	text      string
	timestamp int64
}

// NewQuoteBuilder creates a new quote builder.
func NewQuoteBuilder() QuoteBuilder {
	return &quoteBuilder{
		id:        uuid.New().String(),
		authorId:  uuid.New().String(),
		timestamp: time.Now().UTC().Unix(),
		text:      "",
	}
}

// WithId sets the id of the quote.
func (qb *quoteBuilder) WithId(id string) QuoteBuilder {
	qb.id = id
	return qb
}

// WithAuthorId sets the author id of the quote.
func (qb *quoteBuilder) WithAuthorId(authorId string) QuoteBuilder {
	qb.authorId = authorId
	return qb
}

// WithText sets the text of the quote.
func (qb *quoteBuilder) WithText(text string) QuoteBuilder {
	qb.text = text
	return qb
}

// WithTimestamp sets the timestamp of the quote.
func (qb *quoteBuilder) WithTimestamp(timestamp int64) QuoteBuilder {
	qb.timestamp = timestamp
	return qb
}

// Build builds the quote.
func (qb *quoteBuilder) Build() Quote {
	return Quote{
		ID:        qb.id,
		AuthorID:  qb.authorId,
		Text:      qb.text,
		Timestamp: qb.timestamp,
	}
}
