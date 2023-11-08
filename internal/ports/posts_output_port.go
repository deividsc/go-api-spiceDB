package ports

import (
	"api-spiceDB/internal/core/domain"
	"context"
)

type PostsRepo interface {
	GetById(ctx context.Context, id string) (domain.Post, error)
	Add(ctx context.Context, p domain.Post) error
}
