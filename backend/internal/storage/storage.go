package storage

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

const queryTimeoutDuration = time.Second * 5

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
