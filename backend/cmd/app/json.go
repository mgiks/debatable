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

	if err := writeJSON(w, respCode, envelope{Data: data}); err != nil {
		return fmt.Errorf("failed to write json: %w", err)
	}

	return nil
}

func readJSON(w http.ResponseWriter, r *http.Request, dest any) error {
	max_bytes := 1_048_576 // 1MB
	r.Body = http.MaxBytesReader(w, r.Body, int64(max_bytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(dest)
	if err != nil {
		switch err.(type) {
		case *json.SyntaxError:
			return fmt.Errorf("invalid json")
		case *json.UnmarshalTypeError:
			return fmt.Errorf("inavalid json data")
		default:
			return fmt.Errorf("failed to read json")
		}
	}
	return nil
}

func writeJSONerror(w http.ResponseWriter, respCode int, errMessage string) {
	type envelope struct {
		Error string `json:"error"`
	}
	writeJSON(w, respCode, envelope{
		Error: errMessage,
	})
}

func writeJSON(w http.ResponseWriter, respCode int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(respCode)
	return json.NewEncoder(w).Encode(data)
}
