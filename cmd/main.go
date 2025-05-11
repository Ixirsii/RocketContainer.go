package main

import (
	"RocketContainer.go/graph"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dotenv-org/godotenvvault"
	"github.com/vektah/gqlparser/v2/ast"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	dotenvErr := godotenvvault.Load()
	if dotenvErr != nil {
		logger.Fatal("failed to load .env file", zap.Error(dotenvErr))
	}

	gormLogger := zapgorm2.New(logger)
	gormLogger.SetAsDefault()

	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPass + " dbname=" + dbName + " port=" + dbPort

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
	if dbErr != nil {
		logger.Fatal("failed to connect to database", zap.Error(dbErr))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(
		extension.AutomaticPersistedQuery{
			Cache: lru.New[string](100),
		},
	)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Info("connect to http://localhost:/ for GraphQL playground", zap.String("port", port))
	httpErr := http.ListenAndServe(":"+port, nil)
	if httpErr != nil {
		logger.Fatal("failed to start HTTP server", zap.Error(httpErr))
	}
}
