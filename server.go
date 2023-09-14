package main

import (
	"bingo/graph/generated"
	graph "bingo/graph/resolver"
	"bingo/internal/middleware/auth"
	"bingo/pkg/mongoClient"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const (
	defaultPort     = "8080"
	defaultMongoUrl = "mongodb://localhost:27017/?maxPoolSize=1"
)

// var MongoDB *mongoClient.DB

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	url := os.Getenv("MONGODB_URI")
	if url == "" {
		url = defaultMongoUrl
	}
	client := mongoClient.MongoConnect(url)
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Println("Disconnect MongoDB")
			panic(err)
		}
	}()

	router := chi.NewRouter()
	router.Use(auth.AuthMiddleware())
	router.Use(cors.AllowAll().Handler)
	router.Use(mongoClient.MongoDBMiddleware(client))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/", playground.Handler("Bingo playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for Bingo playground", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(err)
	}
}
