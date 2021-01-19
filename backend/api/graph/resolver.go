package graph

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"TakuroSugahara/plactice-go-project/view/model"
	"context"
	"fmt"
	"math/rand"
)

type Resolver struct {
	users []*model.User
	posts []*model.Post
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	sugahara := &model.User{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Name: "菅原 拓朗",
	}

	post := &model.Post{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Title:  input.Title,
		Author: sugahara,
	}
	r.posts = append(r.posts, post)
	return post, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	sugahara := &model.User{
		ID:   fmt.Sprintf("T%d", rand.Int()),
		Name: "菅原 拓朗",
	}

	if len(r.users) == 0 {
		r.users = append(r.users, sugahara)
	}
	return r.users, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return r.posts, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
