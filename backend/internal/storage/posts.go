package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Post struct {
	Id    int64  `json:"id"`
	Title string `json:"string"`
	Body  string `json:"body"`
}

type postRepository struct {
	db *pgxpool.Pool
}

func (r postRepository) Create(ctx context.Context, post *Post) error {
	return nil
}
