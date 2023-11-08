package ports

import (
	"api-spiceDB/internal/core/domain"
	"context"
)

type PostsInputPort interface {
	GetPosts(ctx context.Context) ([]domain.Post, error)
	AddPost(ctx context.Context, p domain.Post) error
}
