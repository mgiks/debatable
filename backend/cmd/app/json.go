package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (app application) writeJSONresponse(w http.ResponseWriter, respCode int, data any) error {
	type envelope struct {
		Data any `json:"data"`
	}

	err := writeJSON(w, respCode, envelope{Data: data})
	if err != nil {
		return fmt.Errorf("failed to write json: %w", err)
	}

	return nil
}

func writeJSONerror(w http.ResponseWriter, respCode int, errMessage string) {
	type envelope struct {
		Error string `json:"error"`
	}

	w.WriteHeader(respCode)
	writeJSON(w, respCode, envelope{
		Error: errMessage,
	})
}

func writeJSON(w http.ResponseWriter, respCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(respCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}

	return nil
}
