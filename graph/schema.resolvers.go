package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"errors"
	"gogingraphqleg/graph/model"
	"log"

	"github.com/google/uuid"
)

var users []*model.User

func init() {
	log.Println("Init - Users array to be created")
	users = make([]*model.User, 0)
	users = append(users, &model.User{ID: "1", FirstName: "Srinivasa", LastName: "Ramanujam", Dob: "12/27/1887"})
	users = append(users, &model.User{ID: "2", FirstName: "CV", LastName: "Raman", Dob: "11/7/1888"})
	users = append(users, &model.User{ID: "3", FirstName: "Subrahmanyan", LastName: "Chandrasekhar", Dob: "10/19/1910"})
	log.Println("Init - Users array has been created")
}
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	log.Println("Create a new USER")
	uuidValue := uuid.NewString()
	user := &model.User{ID: uuidValue, FirstName: input.FirstName, LastName: input.LastName}
	users = append(users, user)
	return user, nil
}
func (r *mutationResolver) RemoveUser(ctx context.Context, input model.DeleteUser) (*model.User, error) {
	index := -1
	for i, user := range users {
		if user.ID == input.UserID {
			index = i
		}
	}
	if index == -1 {
		return nil, errors.New("Cannot find user you are looking for!")
	}
	user := users[index]
	users = append(users[:index], users[index+1:]...)

	return user, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return users, nil
}
func (r *queryResolver) FindUser(ctx context.Context, id string) (*model.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("Cannot find user you are looking for!")
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return nil, errors.New("Cannot find user you are looking for!")
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
