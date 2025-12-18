package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// http.FileServer 将“设定的根目录”和“请求的URL路径”直接拼到一起去硬盘之中寻找文件。
	fileServer := http.FileServer(http.Dir("./ui/static"))

	// Use the mux.Handle() function...
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
