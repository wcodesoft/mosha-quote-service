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

	// DeleteAuthorQuotes deletes all quotes from an author.
	DeleteAuthorQuotes(authorID string) error

	// GetRandomQuote returns a random quote from the database.
	GetRandomQuote() (data.Quote, error)
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

func (r *repository) authorExist(authorId string) error {
	_, err := r.clientRepository.GetAuthor(authorId)
	return err
}

// AddQuote adds a new quote to the database.
func (r *repository) AddQuote(quote data.Quote) (string, error) {
	if err := r.authorExist(quote.AuthorID); err != nil {
		return "", err
	}
	return r.db.AddQuote(quote)
}

func (r *repository) ListAll() []data.Quote {
	return r.db.ListAll()
}

func (r *repository) UpdateQuote(quote data.Quote) (data.Quote, error) {
	if err := r.authorExist(quote.AuthorID); err != nil {
		return data.Quote{}, err
	}
	return r.db.UpdateQuote(quote)
}

func (r *repository) DeleteQuote(id string) error {
	return r.db.DeleteQuote(id)
}

func (r *repository) GetQuote(id string) (data.Quote, error) {
	return r.db.GetQuote(id)
}

func (r *repository) GetAuthorQuotes(authorID string) []data.Quote {
	return r.db.GetAuthorQuotes(authorID)
}

func (r *repository) DeleteAuthorQuotes(authorID string) error {
	return r.db.DeleteAuthorQuotes(authorID)
}

func (r *repository) GetRandomQuote() (data.Quote, error) {
	return r.db.GetRandomQuote()
}
