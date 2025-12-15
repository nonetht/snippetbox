package main

import (
	"log/slog"
	"net/http"
)

// templateData holds data passed to templates. Extend this struct as you learn new concepts.
type templateData struct {
	// TODO: Add fields for passing dynamic data into templates (e.g. snippets, forms, flash messages).
}

// render is a placeholder to keep handlers lightweight while you wire up templates later.
// Replace this with real template parsing/caching when ready.
func render(logger *slog.Logger, w http.ResponseWriter, r *http.Request, page string, data *templateData) {
	logger.Info("render placeholder", "page", page, "method", r.Method, "uri", r.URL.RequestURI())
	http.Error(w, "render not implemented for "+page, http.StatusNotImplemented)
}

// serverError writes a log entry at Error level (including the request method and URI),
// then sends a generic 500 Internal Server Error response to the user.
func serverError(logger *slog.Logger, w http.ResponseWriter, r *http.Request, err error) {
	logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError sends a specific status code and corresponding description to the user.
func clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
