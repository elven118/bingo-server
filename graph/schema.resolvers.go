package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bingo/graph/generated"
	"bingo/graph/model"
	"bingo/internal/auth"
	"bingo/internal/bingoCard"
	"bingo/pkg/jwt"
	"bingo/pkg/mongoClient"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input model.User) (string, error) {
	log.Printf("login(name: %s, code: %d)", input.Name, input.Code)

	var user User
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
	if user.Name == "" && user.Password == "" {
		// update account info
		user.Name = input.Name
		user.Password = input.Password
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
	} else if user.Name == input.Name && user.Password == input.Password {
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

// BingoCard is the resolver for the bingoCard field.
func (r *queryResolver) BingoCard(ctx context.Context) (*model.BingoCard, error) {
	userId := auth.ForContext(ctx)
	if userId == nil {
		return &model.BingoCard{}, fmt.Errorf("Access Denied")
	}
	log.Printf("User Login with id: %s, name: %s", userId.Id, userId.Name)

	var user User
	// query user
	client := mongoClient.ForContext(ctx)
	coll := client.Database("bingo").Collection("users")
	id, _ := primitive.ObjectIDFromHex(userId.Id)
	err := coll.FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		log.Printf("No document was found with id %s", userId.Id)
		return nil, fmt.Errorf("Please login again.")
	}
	if err != nil {
		fmt.Printf("MongoDB error: %s", err)
		return nil, fmt.Errorf("System Error occur.")
	}

	if user.Numbers != nil {
		return &model.BingoCard{Numbers: user.Numbers}, nil
	}
	return nil, fmt.Errorf("System Error occur.")
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Code     int                `bson:"code"`
	Name     string             `bson:"name,omitempty"`
	Password string             `bson:"password,omitempty"`
	Numbers  []int              `bson:"numbers,omitempty"`
}
