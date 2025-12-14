package main

import (
	"database/sql"
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
	// Define new command-line flag for MySQL DSN string.
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "Mysql data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)) // 也可以使用 NEWJSON...

	// To ... put creating connection pool into openDB function.
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// alse defer a call to db.Close(), so the connection pool is closed before the main() function exits
	defer db.Close()

	app := &application{
		logger: logger,
	}

	logger.Info("starting server", slog.String("addr", ":4000"))

	// 调用 app.routes() 防范让 servemux 包含我们的路径...
	err = http.ListenAndServe(*addr, app.routes())

	// 使用 Error 函数来记录任意的 http.ListenAndServe 函数所返回的 Error 严重级别的信息
	// 随机调用 os.Exit(1) 来终止应用
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Ping 是不是 unix 下的 ping 命令呢？
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
