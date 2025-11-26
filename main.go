package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

/*
GET: 方法是仅返回数据，不修改程序了内容的路由。
POST: 这是改变应用程序内容的路由。
*/

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	// 对snippetView，更新为从请求URL中获取id值。使用此id来从数据库中选择特定代码片段。
	// Extract the value of the id wildcard from the request using r.PathValue()
	// and try to convert it to an integer using the strconv.Atoi() function. If
	// it can't be converted to an integer, or the value is less than 1, we
	// return a 404 page not found response.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Sprintf() function to interpolate the id value with a
	// message, then write it as the HTTP response.
	msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet...")) // 显示存过件新片段的表单
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	// Use w.WriteHeader() method to send 201 status code
	w.WriteHeader(http.StatusCreated)

	// Use w.Write() method to write the response body as normal
	w.Write([]byte("Save a new snippet..."))

}

func main() {
	// Register the two new handler functions and corresponding route patterns with
	// the servemux, in exactly the same way that we did before.
	mux := http.NewServeMux()
	// 单纯的 "/" 只会起到全局捕获效果；写成 "/{$}" 代表着匹配单独的斜杠
	mux.HandleFunc("GET /{$}", home) // HTTP方法区分大小写，并且只能用大写字母。
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	//Create the new route, which is restricted to POST requests
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	log.Print("starting server on :4000")

	// Listen采用：host:port格式，如果省略主机部分，那么将监听所有网络接口
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
