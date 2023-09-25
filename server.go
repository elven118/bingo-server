package main

import (
	"bingo/graph/generated"
	graph "bingo/graph/resolver"
	"bingo/internal/middlewares/auth"
	"bingo/internal/repositories"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	defaultPort     = "8080"
	defaultMongoUrl = "mongodb://localhost:27017/?maxPoolSize=1"
	defaultDatabase = "bingo"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	url := os.Getenv("MONGODB_URI")
	if url == "" {
		url = defaultMongoUrl
	}
	db := os.Getenv("MONGODB_DATABASE")
	if db == "" {
		db = defaultDatabase
	}
	// connect to mongodb
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		log.Printf("Connect to MongoDB error: %s", err)
		panic(err)
	}
	// Ping the primary
	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		log.Printf("Mongo Ping fail: %s", err)
		panic(err)
	}
	log.Printf("Successfully connected and pinged. MongoDB: %s", url)

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Println("Disconnect MongoDB")
			panic(err)
		}
	}()

	router := chi.NewRouter()
	router.Use(auth.AuthenticationMiddleware)
	router.Use(cors.AllowAll().Handler)

	// init resolver
	userRepository := repositories.NewUserRepository(client.Database(db))
	resolver := graph.NewResolver(userRepository)
	// init server
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	router.Handle("/", playground.Handler("Bingo playground", "/graphql"))
	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for Bingo playground", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(err)
	}
}
