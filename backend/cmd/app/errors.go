package main

import "net/http"

func (app application) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.logger.Error("internal server error", "path", r.URL.Path, "method", r.Method, "err", err)
	writeJSONerror(w, http.StatusInternalServerError, "server encountered a problem")
}
