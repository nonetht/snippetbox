package main

import (
	"errors"
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// TODO: Replace this with your preferred logger (configurable formats, levels, destinations).
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// TODO: Add your dependencies (database, template cache, config) and wire them into routes.
	mux := buildRoutes(logger)

	server := &http.Server{
		Addr:    *addr,
		Handler: mux,
	}

	logger.Info("starting server", slog.String("addr", server.Addr))

	// TODO: Consider graceful shutdown with context and signals when you add long-lived resources.
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		logger.Error("server stopped", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

// buildRoutes returns the HTTP routing configuration for the application.
// Add middlewares and new routes here; handlers are defined in separate files.
func buildRoutes(logger *slog.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.Handle("GET /{$}", homeHandler(logger))
	mux.Handle("GET /snippet/view/{id}", snippetViewHandler(logger))
	mux.Handle("GET /snippet/create", snippetCreateHandler(logger))
	mux.Handle("POST /snippet/create", snippetCreatePostHandler(logger))

	return mux
}
