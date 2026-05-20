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
	query := `
		INSERT INTO posts (title, body)
		VALUES ($1, $2) 
		RETURNING id
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	err := r.db.QueryRow(ctx, query, &post.Title, &post.Body).Scan(&post.Id)
	if err != nil {
		return err
	}

	return nil
}
