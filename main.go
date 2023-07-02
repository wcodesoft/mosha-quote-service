package main

import (
	"github.com/charmbracelet/log"
	"github.com/wcodesoft/mosha-quote-service/repository"
	"github.com/wcodesoft/mosha-quote-service/service"
	"os"
	"sync"
)

const (
	defaultHttpPort   = "8280"
	defaultGrpcPort   = "8281"
	QuoteServiceName  = "QuoteService"
	defaultMongoHost  = "mongodb://localhost:27017"
	defaultDatabase   = "mosha"
	authorGrpcAddress = "localhost:8181"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	log.Printf("Starting %s", QuoteServiceName)
	httpPort := getEnv("COMPONENT_PORT", defaultHttpPort)
	mongoHost := getEnv("MONGO_DB_HOST", defaultMongoHost)
	authorServiceAddress := getEnv("AUTHOR_SERVICE_ADDRESS", authorGrpcAddress)
	grpcPort := getEnv("GRPC_PORT", defaultGrpcPort)

	clientsRepository := repository.NewClientRepository(repository.ClientsAddress{
		AuthorServiceAddress: authorServiceAddress,
	})
	database := repository.NewMongoDatabase(mongoHost, defaultDatabase)
	repo := repository.New(database, clientsRepository)
	s := service.New(repo)

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		// Create a new HttpRouter.
		router := service.NewHttpRouter(s, QuoteServiceName)
		router.Start(httpPort)
		wg.Done()
	}()

	go func() {
		grpcRouter := service.NewGrpcRouter(s, QuoteServiceName)
		grpcRouter.Start(grpcPort)
		wg.Done()
	}()

	wg.Wait()
}
