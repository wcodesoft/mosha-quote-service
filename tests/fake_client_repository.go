package tests

import (
	"fmt"
	"github.com/wcodesoft/mosha-author-service/data"
	"github.com/wcodesoft/mosha-quote-service/repository"
)

type fakeClientRepository struct {
	authors []data.Author
}

func (f fakeClientRepository) GetAuthor(id string) (data.Author, error) {
	for _, author := range f.authors {
		if author.ID == id {
			return author, nil
		}
	}
	return data.Author{}, fmt.Errorf("author %q do not exist in database", id)
}

func NewFakeClientRepository() repository.ClientRepository {
	return &fakeClientRepository{
		authors: []data.Author{
			data.NewAuthorBuilder().WithId("123").WithName("William").WithPicUrl("somePic").Build(),
			data.NewAuthorBuilder().WithId("456").WithName("Shakespeare").WithPicUrl("somePic").Build(),
		},
	}
}
