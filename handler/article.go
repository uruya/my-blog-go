package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

func (h *Handler) Article(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		h.getArticle(w, req)
	case http.MethodPost:
		h.createArticle(w, req)
	case http.MethodDelete:
		h.deleteArticle(w, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *Handler) getArticle(w http.ResponseWriter, req *http.Request) {
	queryID := req.URL.Query().Get("id")
	id, err := strconv.Atoi(queryID)
	if err != nil {
		log.Println("クエリパラメーターidのパースに失敗しました:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	a, err := h.article.Get(id)
	if err != nil {
		log.Println("記事の取得に失敗しました:", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	type Article struct {
		ID         int
		Title      string
		Paragraphs []string
		Created    time.Time
	}
	params := struct {
		Title   string
		Article Article
	}{
		Title: "わたしのブログ",
		Article: Article{
			ID:         a.ID,
			Title:      a.Title,
			Created:    a.Created,
			Paragraphs: strings.Split(a.Content, "\n"),
		},
	}
	h.templateArticle.Execute(w, params)
}

func (h *Handler) createArticle(w http.ResponseWriter, req *http.Request) {
	title := req.PostFormValue("title")
	content := req.PostFormValue("content")

	// 長さチェック
	if err := validate(title, content); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	id, err := h.article.Create(title, content)
	if err != nil {
		log.Println("記事の作成に失敗しました:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Location", fmt.Sprintf("/articles?id=%d", id))
	w.WriteHeader(http.StatusSeeOther)
}

func validate(title, content string) error {
	if len(title) == 0 {
		return fmt.Errorf("タイトルを入力してください。")
	}
	if len(content) == 0 {
		return fmt.Errorf("内容を入力してください。")
	}
	if utf8.RuneCountInString(title) > 100 {
		return fmt.Errorf("タイトルは100文字以内におさめてください。")
	}
	if utf8.RuneCountInString(content) > 1000 {
		return fmt.Errorf("内容は1000文字以内におさめてください。")
	}
	return nil
}

func (h *Handler) deleteArticle(res http.ResponseWriter, req *http.Request) {
	// IDを取得
	queryID := req.URL.Query().Get("id")
	id, err := strconv.Atoi(queryID)
	if err != nil {
		log.Println("failed to parse query:", queryID, err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	// 削除を実行
	if err := h.article.Delete(id); err != nil {
		log.Println("failed to delete article:", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}

func (h *Handler) NewArticle(w http.ResponseWriter, r *http.Request) {
	params := struct {
		Title string
	}{
		Title: "わたしのブログ",
	}
	h.templateNewArticle.Execute(w, params)
}
