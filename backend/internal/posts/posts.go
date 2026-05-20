package posts

import (
	"context"

	"github.com/mgiks/debatable/internal/storage"
)

type PostService interface {
	CreatePost(context.Context, *storage.Post) error
}

type postService struct {
	post storage.PostRepository
}

func (s postService) CreatePost(ctx context.Context, post *storage.Post) error {
	return nil
}

func NewPostService(repo storage.PostRepository) PostService {
	return postService{
		post: repo,
	}
}
