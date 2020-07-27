package graph

import (
	"github.com/alexedwards/scs/v2"
	"github.com/kriskelly/dating-app-example/internal/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users          map[string]*model.User
	sessionManager *scs.SessionManager
}

func NewResolver(sessionManager *scs.SessionManager) *Resolver {
	return &Resolver{
		users:          make(map[string]*model.User),
		sessionManager: sessionManager,
	}
}
