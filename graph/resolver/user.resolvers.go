package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.37

import (
	"bingo/graph/model"
	"bingo/internal/middlewares/auth"
	"bingo/internal/models"
	"context"
	"fmt"
)

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, id *string) ([]*model.User, error) {
	user := auth.ForContext(ctx)
	if user == nil || user.Role != string(models.RoleAdmin) {
		return []*model.User{}, fmt.Errorf("Access Denied")
	}

	// query users
	users, err := r.userRepository.FindAllPlayersWithName()
	if err != nil {
		fmt.Printf("MongoDB error: %s", err)
		return nil, fmt.Errorf("System Error occur.")
	}

	return users, nil
}
