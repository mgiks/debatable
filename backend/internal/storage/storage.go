package storage

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Posts() PostRepository
}

type PostRepository interface {
	Create(context.Context, *Post) error
}

type store struct {
	posts postRepository
}

func (s store) Posts() PostRepository {
	return s.posts
}

func NewStore(db *pgxpool.Pool) Store {
	return store{
		posts: postRepository{db: db},
	}
}
