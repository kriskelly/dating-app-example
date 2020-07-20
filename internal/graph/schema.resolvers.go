package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/kriskelly/dating-app-example/internal/graph/generated"
	"github.com/kriskelly/dating-app-example/internal/model"
)

func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.User, error) {
	user := r.users[email]
	if user == nil {
		return nil, errors.New("Invalid email")
	}
	if user.Password != password {
		return nil, errors.New("Invalid password")
	}
	r.currentUser = user
	return user, nil
}

func (r *mutationResolver) Signup(ctx context.Context, input model.NewUser) (*model.User, error) {
	newUser := &model.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}
	id := uuid.New().String()
	newUser.ID = id
	r.users[newUser.Email] = newUser
	return newUser, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	user := r.currentUser
	if user == nil {
		return nil, errors.New("No current user")
	}
	return user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
