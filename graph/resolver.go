package graph

//go:generate go run github.com/99designs/gqlgen generate

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/marianavif/gqlgen-todos/graph/generated"
	"github.com/marianavif/gqlgen-todos/graph/model"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	todos []*model.Todo
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Todo() generated.TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

type todoResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	todo := &model.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", randNumber),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}
