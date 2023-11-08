package application

import (
	"api-spiceDB/internal/core/domain"
	"api-spiceDB/internal/ports"
	"context"
	"errors"
)

type GetPostByIdHandler interface {
	GetPost(ctx context.Context, id string) (domain.Post, error)
}

type GetPostByIdHandlerImpl struct {
	repo ports.PostsRepo
}

func NewGetPostsHandlerImp(repo ports.PostsRepo) (*GetPostByIdHandlerImpl, error) {
	if repo == nil {
		return nil, errors.New("repository cannot be nil")
	}
	return &GetPostByIdHandlerImpl{
		repo: repo,
	}, nil
}

func (h *GetPostByIdHandlerImpl) GetPost(ctx context.Context, id string) (domain.Post, error) {
	return h.repo.GetById(ctx, id)
}
