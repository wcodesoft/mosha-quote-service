package service

import (
	"github.com/wcodesoft/mosha-quote-service/data"
	"github.com/wcodesoft/mosha-quote-service/repository"
)

// Service represents the service interface.
type Service interface {
	// CreateQuote registers a new Quote in the database.
	CreateQuote(quote data.Quote) (string, error)

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
}

type service struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &service{
		repo: repo,
	}
}

// CreateQuote registers a new Quote in the database.
func (s *service) CreateQuote(quote data.Quote) (string, error) {
	return s.repo.AddQuote(quote)
}

// ListAll returns all quotes in the database.
func (s *service) ListAll() []data.Quote {
	return s.repo.ListAll()
}

// UpdateQuote updates a quote in the database.
func (s *service) UpdateQuote(quote data.Quote) (data.Quote, error) {
	return s.repo.UpdateQuote(quote)
}

// DeleteQuote deletes a quote from the database.
func (s *service) DeleteQuote(id string) error {
	return s.repo.DeleteQuote(id)
}

// GetQuote returns a quote from the database.
func (s *service) GetQuote(id string) (data.Quote, error) {
	return s.repo.GetQuote(id)
}

// GetAuthorQuotes returns all quotes from an author.
func (s *service) GetAuthorQuotes(authorID string) []data.Quote {
	return s.repo.GetAuthorQuotes(authorID)
}
