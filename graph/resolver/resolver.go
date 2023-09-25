package graph

import "bingo/internal/repositories"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userRepository *repositories.UserRepository
}

func NewResolver(userRepository *repositories.UserRepository) *Resolver {
	return &Resolver{
		userRepository: userRepository,
	}
}
