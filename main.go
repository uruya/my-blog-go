package main

import (
	"embed"
	"html/template"
	"log"
	"my-blog/article"
	"my-blog/db"
	"my-blog/handler"
	"net/http"
	"os"
)

//go:embed assets
var assets embed.FS

func main() {
	dbconfig := &db.Config{
		Host:     getenv("MY_BLOG_DB_HOST", "localhost"),
		Port:     getenv("MY_BLOG_DB_PORT", "5432"),
		User:     getenv("MY_BLOG_DB_USER", "postgres"),
		Password: getenv("MY_BLOG_DB_PASSWORD", "postgres"),
		Database: getenv("MY_BLOG_DB_DATABASE", "my_blog"),
		SSL:      getenv("MY_BLOG_DB_MODE", "disable"),
	}
	d, err := db.New(dbconfig)
	if err != nil {
		log.Fatal(err)
	}
	a := article.New(d)
	h := handler.New(
		template.Must(template.ParseFS(assets, "assets/index.html")),
		template.Must(template.ParseFS(assets, "assets/article.html")),
		template.Must(template.ParseFS(assets, "assets/new.html")),
		a,
	)
	http.HandleFunc("/", h.Index)
	http.HandleFunc("/articles", h.Article)
	http.HandleFunc("/articles/new", h.NewArticle)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getenv(key, defaultValue string) string {
	if env := os.Getenv(key); env != "" {
		return env
	}
	return defaultValue
}
