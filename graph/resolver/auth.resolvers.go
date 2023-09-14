package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bingo/graph/generated"
	"bingo/graph/model"
	"bingo/internal/app/bingoCard"
	"bingo/internal/models"
	"bingo/pkg/jwt"
	"bingo/pkg/mongoClient"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.User) (string, error) {
	log.Printf("login(name: %s, code: %d)", input.Name, input.Code)

	var user models.User
	// query user
	client := mongoClient.ForContext(ctx)
	coll := client.Database("bingo").Collection("users")
	err := coll.FindOne(context.Background(), bson.D{{Key: "code", Value: input.Code}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		log.Printf("Code not found %s", input.Code)
		return "", fmt.Errorf("Code not correct")
	}
	if err != nil {
		fmt.Printf("MongoDB error: %s", err)
		return "", fmt.Errorf("System Error occur.")
	}

	// check account name password
	if user.Name == "" {
		// update account info
		user.Name = input.Name
		user.Numbers = bingoCard.GenBingoCard()
		userReplace, err := coll.ReplaceOne(context.Background(), bson.D{{Key: "_id", Value: user.Id}}, user)
		if err != nil {
			log.Println(err)
			panic(err)
		}
		log.Printf("User replace result: %v", userReplace)

		token, err := jwt.GenToken(user.Id.Hex(), input.Name)
		if err != nil {
			log.Printf("Gen Token Error: %v", err)
			return "", err
		} else {
			return token, nil
		}
	} else if user.Name == input.Name && user.Code == input.Code {
		token, err := jwt.GenToken(user.Id.Hex(), input.Name)
		if err != nil {
			log.Printf("Gen Token Error: %v", err)
			return "", err
		} else {
			return token, nil
		}
	}

	return "", fmt.Errorf("Username or Password not correct")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
