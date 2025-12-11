package main

import (
	"net/http"
)

// The serverError helper writes a log entry at Error level (including the request
// method and URI as attributes), then sends a generic 500 Internal Server Error
// response to the user.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	// log entry at Error level, including the request method and URL as attributes
	app.logger.Error(err.Error(), "method", method, "uri", uri)
	// send a generic 500 internal Server Error response
	// http.StatusText() 函数，可以返回指定HTTP状态码的可读文本描述，比如：http.StatusText(400) 返回 "Bad Request"
	// http.StatusText(500) 返回 "Internal Server Error"
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
// to the user. We'll use this later in the book to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}
