package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "http service address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	mux := http.NewServeMux()

	// 将请求 URL 映射到 ./ui/static/ 目录下文件读出来并返回
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// 1.受到 GET /static/ 请求，随即交给 handler 处理
	// 2. 将请求的开头 "/static" 去掉，比如请求是 /static/css/main.css，最后变成了 ./ui/static/css/main.csss
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	//logger.Info("request receivced", "method", "GET", "path", "/")

	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	logger.Info("starting server", "addr", addr)

	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error()) // Use Error() method to log any **error message**
	os.Exit(1)
}
