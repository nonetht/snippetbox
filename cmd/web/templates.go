package main

import (
	"html/template"
	"path/filepath"

	"snippetbox.yang.net/internal/models"
)

// 构建一个结构体，即 templateData 类型，去承载任意动态数据
// 此时我们仅包含一个字段，但是后续会添加更多内容。
type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initialize a new map
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		// Base return the last element of page, like "./ui/html/pages/home.tmpl" -> "home.tmpl"
		name := filepath.Base(page)

		ts, err := template.ParseFiles("./ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}

		// Call ...
		ts, err = ts.ParseGlob("./ui/html/base/*.tmpl")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map as normal
		cache[name] = ts
	}
	return cache, nil
}
