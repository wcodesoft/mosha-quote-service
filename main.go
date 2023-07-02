package main

import (
	"context"
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/wcodesoft/mosha-quote-service/repository"
	"github.com/wcodesoft/mosha-quote-service/service"
	"net/http"
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

// InterceptorLogger adapts slog logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l *log.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		switch lvl {
		case logging.LevelDebug:
			l.Debugf(msg, fields)
		case logging.LevelInfo:
			l.Infof(msg, fields)
		case logging.LevelWarn:
			l.Warnf(msg, fields)
		case logging.LevelError:
			l.Errorf(msg, fields)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
	})
}
func main() {
	log.Printf("Starting %s", QuoteServiceName)
	port := getEnv("COMPONENT_PORT", defaultHttpPort)
	mongoHost := getEnv("MONGO_DB_HOST", defaultMongoHost)
	authorServiceAddress := getEnv("AUTHOR_SERVICE_ADDRESS", authorGrpcAddress)

	clientsRepository := repository.NewClientRepository(repository.ClientsAddress{
		AuthorServiceAddress: authorServiceAddress,
	})
	database := repository.NewMongoDatabase(mongoHost, defaultDatabase)
	repo := repository.New(database, clientsRepository)
	s := service.New(repo)

	wg := new(sync.WaitGroup)

	wg.Add(1)

	go func() {
		log.Infof("Starting %s http on %s", QuoteServiceName, port)
		// Create a new HttpRouter.
		router := service.NewHttpRouter(s)
		if err := http.ListenAndServe(fmt.Sprintf(":%s", port), router.MakeHandler()); err != nil {
			log.Fatalf("Unable to start service %q: %s", QuoteServiceName, err)
		}
		wg.Done()
	}()

	wg.Wait()
}
