package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql" // 并未使用 mysql 包中的任何内容，只是需要运行驱动程序的 init() 函数
	// {your-module-path}/internal/models
	"snippetbox.yang.net/internal/models"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

// 解析应用程序运行时配置设置，建立处理器依赖关系，运行HTTP服务器
func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Define new command-line flag for MySQL DSN string.
	dsn := flag.String("dsn", "web:password@/snippetbox?parseTime=true", "Mysql data source name")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil)) // 也可以使用 NEWJSON...

	// To ... put creating connection pool into openDB function.
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	// also defer a call to db.Close(), so the connection pool is closed before the main() function exits
	defer db.Close()

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	logger.Info("starting server", slog.String("addr", ":4000"))

	err = http.ListenAndServe(*addr, app.routes())

	// 使用 Error 函数来记录任意的 http.ListenAndServe 函数所返回的 Error 严重级别的信息
	// 随机调用 os.Exit(1) 来终止应用
	logger.Error(err.Error())
	os.Exit(1)
}

// 实际上并不会创建链接，所做的仅仅是初始化连接池，以供后续调用
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// db.Ping() 创建一个链接，并检查是否有任何错误
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
