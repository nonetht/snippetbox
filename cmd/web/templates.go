package main

import "snippetbox.yang.net/internal/models"

// 构建一个结构体，即 templateData 类型，去承载任意动态数据
// 此时我们仅包含一个字段，但是后续会添加更多内容。
type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
