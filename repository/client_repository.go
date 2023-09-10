package repository

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/wcodesoft/mosha-author-service/data"
	mgrpc "github.com/wcodesoft/mosha-service-common/grpc"
	pb "github.com/wcodesoft/mosha-service-common/protos/authorservice"
)

type ClientRepository interface {
	GetAuthor(id string) (data.Author, error)
}
type clientRepository struct {
	client pb.AuthorServiceClient
}

func (c *clientRepository) GetAuthor(id string) (data.Author, error) {
	author, err := c.client.GetAuthor(context.Background(), &pb.GetAuthorRequest{Id: id})
	if err != nil {
		log.Errorf("could not get author with id: %s", id)
		return data.Author{}, err
	}
	return data.Author{
		ID:     author.Id,
		Name:   author.Name,
		PicURL: author.PicUrl,
	}, nil
}

// NewClientRepository creates a new client repository.
func NewClientRepository(clientInfo mgrpc.ClientInfo) (ClientRepository, error) {
	conn, err := clientInfo.NewClientConnection()
	if err != nil {
		return nil, err
	}
	client := pb.NewAuthorServiceClient(conn)
	return &clientRepository{
		client: client,
	}, nil
}
