package main

import (
	"github.com/charmbracelet/log"
	"github.com/wcodesoft/mosha-quote-service/repository"
	"github.com/wcodesoft/mosha-quote-service/service"
	mdb "github.com/wcodesoft/mosha-service-common/database"
	mgrpc "github.com/wcodesoft/mosha-service-common/grpc"
	mhttp "github.com/wcodesoft/mosha-service-common/http"
	"github.com/wcodesoft/mosha-service-common/tracing"
	"os"
	"strconv"
	"sync"
)

const (
	defaultHttpPort       = "8280"
	defaultGrpcPort       = "8281"
	QuoteServiceName      = "QuoteService"
	defaultMongoHost      = "mongodb://localhost:27017"
	defaultDatabase       = "mosha"
	authorGrpcAddress     = "localhost:8181"
	defaultReleaseVersion = "dev"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	log.Printf("Starting %s", QuoteServiceName)
	httpPort := getEnv("HTTP_PORT", defaultHttpPort)
	mongoHost := getEnv("MONGO_DB_HOST", defaultMongoHost)
	authorServiceAddress := getEnv("AUTHOR_SERVICE_ADDRESS", authorGrpcAddress)
	grpcPort := getEnv("GRPC_PORT", defaultGrpcPort)
	releaseVersion := getEnv("RELEASE_VERSION", defaultReleaseVersion)

	sentryDsn := getEnv("SENTRY_DSN", "__DSN__")
	sentrySampleRate, err := strconv.ParseFloat(getEnv("SENTRY_SAMPLE_RATE", "1.0"), 64)
	if err != nil {
		sentrySampleRate = 1.0
	}

	if err := tracing.SetupSentry(sentryDsn, releaseVersion, QuoteServiceName, sentrySampleRate); err != nil {
		log.Error("unable to setup sentry: ", err)
	}

	quoteGrpcClientInfo := mgrpc.ClientInfo{
		Name:    "AuthorService",
		Address: authorServiceAddress,
	}
	clientsRepository, err := repository.NewClientRepository(quoteGrpcClientInfo)
	if err != nil {
		log.Fatal(err)
	}

	mongoClient, err := mdb.NewMongoClient(mongoHost)
	if err != nil {
		log.Fatal(err)
	}
	connection := mdb.NewMongoConnection(mongoClient, defaultDatabase, "quotes")
	database := repository.NewMongoDatabase(connection)
	repo := repository.New(database, clientsRepository)
	s := service.New(repo)

	wg := new(sync.WaitGroup)

	wg.Add(2)

	go func() {
		// Create a new QuoteService.
		hs := service.QuoteService{
			Service: s,
			Port:    httpPort,
			Name:    QuoteServiceName,
		}
		err := mhttp.StartHttpService(&hs)
		if err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	go func() {
		grpcRouter := service.NewGrpcRouter(s, QuoteServiceName)
		if err := grpcRouter.Start(grpcPort); err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	wg.Wait()
}
