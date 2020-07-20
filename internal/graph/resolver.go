package graph

import "github.com/kriskelly/dating-app-example/internal/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users       map[string]*model.User
	currentUser *model.User
}

func NewResolver() *Resolver {
	return &Resolver{
		users: make(map[string]*model.User),
	}
}
