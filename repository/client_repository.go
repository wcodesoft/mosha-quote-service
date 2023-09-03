package repository

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/wcodesoft/mosha-author-service/data"
	pb "github.com/wcodesoft/mosha-service-common/protos/authorservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientRepository interface {
	GetAuthor(id string) (data.Author, error)
}

type ClientsAddress struct {
	AuthorServiceAddress string
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
func NewClientRepository(address ClientsAddress) ClientRepository {
	conn, err := grpc.Dial(address.AuthorServiceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Errorf("Could not connect to AuthorService at: %s", address.AuthorServiceAddress)
		panic(err)
	}
	client := pb.NewAuthorServiceClient(conn)
	log.Infof("Connected to AuthorService at: %s", address.AuthorServiceAddress)
	return &clientRepository{
		client: client,
	}
}
