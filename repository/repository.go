package repository

import "github.com/wcodesoft/mosha-quote-service/data"

// Repository represents the repository interface.
type Repository interface {
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
}

type repository struct {
	db               Database
	clientRepository ClientRepository
}

// New creates a new repository.
func New(db Database, clientRepository ClientRepository) Repository {
	return &repository{
		db:               db,
		clientRepository: clientRepository,
	}
}

func (s repository) authorExist(authorId string) error {
	_, err := s.clientRepository.GetAuthor(authorId)
	return err
}

// AddQuote adds a new quote to the database.
func (s repository) AddQuote(quote data.Quote) (string, error) {
	if err := s.authorExist(quote.AuthorID); err != nil {
		return "", err
	}
	return s.db.AddQuote(quote)
}

func (s repository) ListAll() []data.Quote {
	return s.db.ListAll()
}

func (s repository) UpdateQuote(quote data.Quote) (data.Quote, error) {
	if err := s.authorExist(quote.AuthorID); err != nil {
		return data.Quote{}, err
	}
	return s.db.UpdateQuote(quote)
}

func (s repository) DeleteQuote(id string) error {
	return s.db.DeleteQuote(id)
}

func (s repository) GetQuote(id string) (data.Quote, error) {
	return s.db.GetQuote(id)
}

func (s repository) GetAuthorQuotes(authorID string) []data.Quote {
	return s.db.GetAuthorQuotes(authorID)
}
