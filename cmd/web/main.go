package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// Difine a new command-line flag with name 'addr', a default value of ":4000"
	// usage: 'xxx': short help text explaining what the flag controls
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Use flag.Parse() function to *parse* the command-line flag.
	// Read in the command-line flag value and assigns it to the addr variable
	flag.Parse()

	// Use the logger.New() function to initialize a new structed logger, which writes
	// to the standard out stream and uses default settings
	//logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	// 也可以通过 NewJSONHandler 创建一个将日志条目以JSON对象写入的处理器
	// 通过slog.New() 创建的自定义日志记录器是并发安全的，无须担心竞态条件
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" 文件夹
	// The given path is relative to the project dir root
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server
	// StripPrefix Remove the given prefix from the given URL
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Register the other application routes as normal
	// http.HandlerFunc 适配器工作原理就是，自动为函数添加 `ServeHTTP`方法。执行时，自动调用原始函数内部代码。
	// 将普通函数强制转换为满足 `http.Handler` 接口的类型
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// 使用Info函数，来记录服务器启动记录。
	logger.Info("starting server", slog.String("addr", ":4000"))

	// 尽管 ListAndServe 第二个接收的参数是 http.Handler，但我们一直使用的 servemux 来传递参数
	// 这是因为 servemux 也实现了 ServeHTTP方法，同样满足 Handler 接口
	err := http.ListenAndServe(*addr, mux)

	// And, use Error() method 记录 http.ListenAndServe 的错误信息。
	// 随后调用 os.Exit(1)来终止程序，并伴有终止符1。
	logger.Error(err.Error())
	os.Exit(1)
}
