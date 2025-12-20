package main

import "net/http"

func (app *application) serveError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		url    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "path", url)
	// http.StatusText() 返回指定 HTTP 状态码的可读文本描述
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, r *http.Request, status int) {
	http.Error(w, http.StatusText(status), status)
}
