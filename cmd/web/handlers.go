package main

import (
	"log/slog"
	"net/http"
	"strconv"
)

// homeHandler renders the landing page. Extend templateData to pass dynamic values.
func homeHandler(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Populate template data when templates are wired up.
		render(logger, w, r, "home.tmpl", nil)
	})
}

// snippetViewHandler shows a single snippet. Replace stub with real lookup logic.
func snippetViewHandler(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil || id < 1 {
			clientError(w, http.StatusNotFound)
			return
		}

		// TODO: Load a snippet by ID and pass data into the template.
		render(logger, w, r, "snippet_view.tmpl", nil)
	})
}

// snippetCreateHandler displays a form for creating a snippet.
func snippetCreateHandler(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Show the form with any default values or CSRF token.
		render(logger, w, r, "snippet_create.tmpl", nil)
	})
}

// snippetCreatePostHandler accepts form submissions to create a snippet.
func snippetCreatePostHandler(logger *slog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: Parse form data, validate input, and persist to storage.
		http.Error(w, "snippetCreatePost not implemented yet", http.StatusNotImplemented)
	})
}
