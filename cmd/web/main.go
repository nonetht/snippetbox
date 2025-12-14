package main

import (
	"flag"
	"log/slog" // 包含内容：时间戳，日志条目严重级别，日志消息，可选数量键值对
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

// 解析应用程序运行时配置设置，建立处理器依赖关系，运行HTTP服务器
func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)) // 也可以使用 NEWJSON...

	app := &application{
		logger: logger,
	}

	logger.Info("starting server", slog.String("addr", ":4000"))

	// 调用 app.routes() 防范让 servemux 包含我们的路径...
	err := http.ListenAndServe(*addr, app.routes())

	// 使用 Error 函数来记录任意的 http.ListenAndServe 函数所返回的 Error 严重级别的信息
	// 随机调用 os.Exit(1) 来终止应用
	logger.Error(err.Error())
	os.Exit(1)
}
