package repository

import (
	"fmt"
	"github.com/wcodesoft/mosha-quote-service/data"
	"math/rand"
)

type inMemoryDatabase struct {
	storage map[string]data.Quote
}

func (db *inMemoryDatabase) AddQuote(quote data.Quote) (string, error) {
	if _, ok := db.storage[quote.ID]; ok {
		return "", fmt.Errorf("quote %q already exist in database", quote.ID)
	}
	db.storage[quote.ID] = quote
	return quote.ID, nil
}

func (db *inMemoryDatabase) ListAll() []data.Quote {
	var quotes []data.Quote
	for _, v := range db.storage {
		quotes = append(quotes, v)
	}
	return quotes
}

func (db *inMemoryDatabase) UpdateQuote(quote data.Quote) (data.Quote, error) {
	if _, ok := db.storage[quote.ID]; !ok {
		return data.Quote{}, fmt.Errorf("quote %q do not exist in database", quote.ID)
	}
	db.storage[quote.ID] = quote
	return db.storage[quote.ID], nil
}

func (db *inMemoryDatabase) DeleteQuote(id string) error {
	if _, ok := db.storage[id]; !ok {
		return fmt.Errorf("quote %q do not exist in database", id)
	}
	delete(db.storage, id)
	return nil
}

func (db *inMemoryDatabase) GetQuote(id string) (data.Quote, error) {
	if _, ok := db.storage[id]; !ok {
		return data.Quote{}, fmt.Errorf("quote %q do not exist in database", id)
	}
	return db.storage[id], nil
}

func (db *inMemoryDatabase) GetAuthorQuotes(authorID string) []data.Quote {
	var quotes []data.Quote
	for _, v := range db.storage {
		if v.AuthorID == authorID {
			quotes = append(quotes, v)
		}
	}
	return quotes
}

func (db *inMemoryDatabase) DeleteAuthorQuotes(authorID string) error {
	for _, v := range db.storage {
		if v.AuthorID == authorID {
			delete(db.storage, v.ID)
		}
	}
	return nil
}

func (db *inMemoryDatabase) GetRandomQuote() (data.Quote, error) {
	var quotes []data.Quote
	for _, v := range db.storage {
		quotes = append(quotes, v)
	}

	if len(quotes) == 0 {
		return data.Quote{}, fmt.Errorf("no quotes in database")
	}
	return quotes[rand.Intn(len(quotes))], nil
}

func NewInMemoryDatabase() Database {
	return &inMemoryDatabase{
		storage: make(map[string]data.Quote),
	}
}
