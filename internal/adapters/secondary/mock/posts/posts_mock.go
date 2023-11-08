package mock

import (
	"api-spiceDB/internal/core/domain"
	"context"
)

type PostsRepoMock struct {
	posts []domain.Post
}

func NewPostsRepoMock() *PostsRepoMock {
	return &PostsRepoMock{
		posts: []domain.Post{},
	}
}

func (m *PostsRepoMock) Get(_ context.Context) ([]domain.Post, error) {
	return m.posts, nil
}

func (m *PostsRepoMock) Add(_ context.Context, p domain.Post) error {
	m.posts = append(m.posts, p)
	return nil
}
