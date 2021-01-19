package graph

import (
	"TakuroSugahara/plactice-go-project/view/model"
	"context"
)

type postResolver struct{ *Resolver }

func (p *postResolver) Author(ctx context.Context, obj *model.Post) (*model.User, error) {
	return &model.User{ID: obj.AuthorID, Name: "User " + obj.AuthorID}, nil
}
