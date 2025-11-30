package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// Difine a new command-line flag with name 'addr', a default value of ":4000"
	// usage: 'xxx': short help text explaining what the flag controls
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Use flag.Parse() function to parse the command-line flag.
	// Read in the command-line flag value and assigns it to the addr variable
	flag.Parse()

	// Use the logger.New() function to initialize a new structed logger, which writes
	// to the standard out stream and uses default settings
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory
	// Note that the path given to the http.Dir function is relative to the project
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server
	// StripPrefix Remove the given prefix from the given URL
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register the other application routes as normal
	// http.HandlerFunc 适配器工作原理就是，自动为函数添加 `ServeHTTP`方法。执行时，自动调用原始函数内部代码。
	// 将普通函数强制转换为满足 `http.Handler` 接口的类型
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// 使用Info函数，来记录服务器启动记录。
	logger.Info("starting server", slog.Any("addr", ":4000"))

	err := http.ListenAndServe(*addr, mux)

	// And, use Error() method 记录 http.ListenAndServe 的错误信息。
	// 随后调用 os.Exit(1)来终止程序，并伴有终止符1。
	logger.Error(err.Error())
	os.Exit(1)
}
