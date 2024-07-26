package handler

import (
	_ "embed"
	"log"
	"my-blog/article"
	"net/http"
	"time"
)

type Summary struct {
	ID      int
	Title   string
	Summary string
	Created time.Time
}

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	articles, err := h.article.GetAll()
	if err != nil {
		log.Println("記事一覧の取得に失敗しました:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	params := struct {
		Title    string
		Summaries []Summary
	}{
		Title:    "わたしのブログ",
		
		Summaries: toSummaries(articles),
	}
	h.templateIndex.Execute(w, params)
}

func toSummaries(articles []article.Article) []Summary {
	summaries := make([]Summary, 0, len(articles))
	for _, a := range articles {
		summaries = append(
			summaries,
			Summary{
				ID:      a.ID,
				Title:   a.Title,
				Summary: summarize(a.Content, 90),
				Created: a.Created,
			},
		)
	}
	return summaries
}

func summarize(s string, length int) string {
	r := []rune(s)
	if len(r) <= length {
		return s
	}
	return string(r[:length]) + "..."
}
