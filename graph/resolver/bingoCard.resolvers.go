package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bingo/graph/generated"
	"bingo/graph/model"
	"bingo/internal/middleware/auth"
	"bingo/internal/models"
	"bingo/pkg/mongoClient"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// BingoCard is the resolver for the bingoCard field.
func (r *queryResolver) BingoCard(ctx context.Context) (*model.BingoCard, error) {
	userId := auth.ForContext(ctx)
	if userId == nil {
		return &model.BingoCard{}, fmt.Errorf("Access Denied")
	}
	log.Printf("User Login with id: %s, name: %s", userId.Id, userId.Name)

	var user models.User
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

// ValidateCard is the resolver for the validateCard field.
func (r *queryResolver) ValidateCard(ctx context.Context, id string) (*model.ValidateResult, error) {
	panic(fmt.Errorf("not implemented: ValidateCard - validateCard"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
