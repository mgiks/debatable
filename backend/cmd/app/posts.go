package main

import (
	"fmt"
	"net/http"

	"github.com/mgiks/debatable/internal/storage"
)

type createPostPayload struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (app application) createPostHandler(w http.ResponseWriter, r *http.Request) {
	var payload createPostPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	post := storage.Post{
		Title: payload.Title,
		Body:  payload.Body,
	}
	if err := app.postService.CreatePost(r.Context(), &post); err != nil {
		app.internalServerError(w, r, fmt.Errorf("failed to create post: %w", err))
		return
	}

	if err := app.writeJSONresponse(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
	}
}
