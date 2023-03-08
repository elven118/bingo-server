package mongoClient

import (
	"context"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var userCtxKey = &contextKey{"db"}

type contextKey struct {
	id string
}

func MongoConnect(url string) *mongo.Client {
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

	return client
}

func MongoDBMiddleware(client *mongo.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, client)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *mongo.Client {
	raw, ok := ctx.Value(userCtxKey).(*mongo.Client)
	if !ok {
		return nil
		// , errors.New("could not get database connection pool from context")
	}
	return raw
}
