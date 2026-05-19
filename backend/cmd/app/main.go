package main

import (
	"log/slog"
)

const version = "0.0.1"

func main() {
	config := config{
		env:  "dev",
		port: ":8080",
		server: serverConfig{
			writeTimeout: "30s",
			readTimeout:  "10s",
			idleTimeout:  "1m",
		},
	}

	logger := slog.Default()

	app := application{
		config: config,
		logger: logger,
	}

	mux := app.mount()

	err := app.run(mux)
	if err != nil {
		logger.Error("app failed to run", "err", err)
	}
}
