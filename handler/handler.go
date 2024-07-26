package handler

import (
	"html/template"
	"my-blog/article"
)

type Handler struct {
	templateIndex      *template.Template
	templateArticle    *template.Template
	templateNewArticle *template.Template
	article            *article.Service
}

func New(
	templateIndex *template.Template,
	templateArticle *template.Template,
	templateNewArticle *template.Template,
	article *article.Service,
) *Handler {
	return &Handler{
		templateIndex:      templateIndex,
		templateArticle:    templateArticle,
		templateNewArticle: templateNewArticle,
		article:            article,
	}
}
