package main

import (
	"context"

	"github.com/mgiks/debatable/internal/db"
	"github.com/mgiks/debatable/internal/env"
	"github.com/mgiks/debatable/internal/logger"
	"github.com/mgiks/debatable/internal/posts"
	"github.com/mgiks/debatable/internal/storage"
)

const version = "0.0.1"

func main() {
	config := config{
		env:  "development",
		port: ":8080",
		server: serverConfig{
			writeTimeout: env.GetString("SERVER_WRITE_TIMEOUT", "30s"),
			readTimeout:  env.GetString("SERVER_READ_TIMEOUT", "10s"),
			idleTimeout:  env.GetString("SERVER_IDLE_TIMEOUT", "1m"),
		},
		db: dbConfig{
			url:             env.GetString("DB_URL", "postgres://admin:adminpassword@localhost:5433/typo-typer"),
			maxConns:        env.GetInt32("DB_MAX_CONNS", 35),
			minIdleConns:    env.GetInt32("DB_MIN_IDLE_CONNS", 5),
			maxConnIdleTime: env.GetString("DB_MAX_CONN_IDLE_TIME", "15m"),
		},
	}

	logger := logger.NewZapLogger(config.env)

	db, err := db.New(context.Background(), config.db.url, config.db.maxConns, config.db.minIdleConns, config.db.maxConnIdleTime)
	if err != nil {
		logger.Error("database initialization failed", "err", err)
		return
	}

	store := storage.NewStore(db)
	postService := posts.NewPostService(store.Posts())

	app := application{
		config:      config,
		logger:      logger,
		postService: postService,
	}

	mux := app.mount()

	err = app.run(mux)
	if err != nil {
		logger.Error("app failed to run", "err", err)
	}
}
