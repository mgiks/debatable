package main

import (
	"fmt"
	"net/http"
)

type healthCheckPayload struct {
	Status  string
	Env     string
	Version string
}

func (app application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	payload := healthCheckPayload{
		Status:  "ok",
		Env:     app.config.env,
		Version: version,
	}

	if err := app.writeJSONresponse(w, http.StatusOK, payload); err != nil {
		app.internalServerError(w, r, fmt.Errorf("failed to write json response: %w", err))
	}
}
