package handler

import (
	"my-blog/test"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestArticle(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want int
	}{
		{"success", "1", http.StatusOK},
		{"bad request", "abc", http.StatusBadRequest},
		{"not found", "-1", http.StatusNotFound},
	}

	h, close := newHandler(t)
	defer close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/articles?id="+tt.id, nil)
			rec := httptest.NewRecorder()
			h.Article(rec, req)

			if rec.Code != tt.want {
				t.Fatalf("want: %d, got: %d", tt.want, rec.Code)
			}

			if rec.Code == http.StatusOK {
				// 記事のタイトルテスト
				if !strings.Contains(rec.Body.String(), `<h2 class="text-start py-1">`) {
					t.Fatal("タイトルがありません。")
				}

				// 記事の内容テスト
				if !strings.Contains(rec.Body.String(), `<p class="my-4">`) {
					t.Fatal("記事の中身がありません。")
				}
			}
		})
	}
}

func TestNewArticle(t *testing.T) {
	h, close := newHandler(t)
	defer close()

	req := httptest.NewRequest(http.MethodGet, "/articles/new", nil)
	rec := httptest.NewRecorder()
	h.NewArticle(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatal("status code is not 200:", rec.Code)
	}

	// タイトルのテスト
	if !strings.Contains(rec.Body.String(), `<h1 class="fs-2">わたしのブログ</h1>`) {
		t.Fatal("title is missing")
	}

	// サブタイトルのテスト
	if !strings.Contains(rec.Body.String(), `<h2 class="my-3">新規作成</h2>`) {
		t.Fatal("subtitle is missing")
	}
	if !strings.Contains(rec.Body.String(), `<form class="my-3" action="/articles" method="post">`) {
		t.Fatal("form is missing")
	}
}

func TestCreateArticle(t *testing.T) {
	h, close := newHandler(t)
	defer close()

	var (
		tooLongTitle = strings.Repeat("あ", 101)
		tooLongContent = strings.Repeat("い", 1001)
		validTitle = strings.Repeat("う", 100)
		validContent = strings.Repeat("え", 1000)
	)
	tests := []struct {
		name	string
		title string
		content string
		want int
	} {
		{"too long title", tooLongTitle, validContent, http.StatusBadRequest},
		{"too long content", validTitle, tooLongContent, http.StatusBadRequest},
		{"success", validTitle, validContent, http.StatusSeeOther},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// リクエストボディ
			form := url.Values{
				"title": {tt.title},
				"content": {tt.content},
			}
			body := strings.NewReader(form.Encode())

			// リクエストの作成
			req := httptest.NewRequest(http.MethodPost, "/articles", body)
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

			// 実行
			rec := httptest.NewRecorder()
			h.Article(rec, req)

			// ステータスコードのテスト
			test.Eq(t, tt.want, rec.Code)

			// レスポンスのテスト
			if tt.want == http.StatusSeeOther {
				test.Eq(t, "/articles?id=10000", rec.Header().Get("Location"))
			}
		})
	}
}

func TestDeleteArticle(t *testing.T) {
	tests := []struct {
		name	string
		id string
		want int
	} {
		{"success", "1", http.StatusNoContent},
		{"bad request", "abc", http.StatusBadRequest},
		{"do nothing", "-1", http.StatusNoContent},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T)  {
			h, close := newHandler(t)
			defer close()

			req := httptest.NewRequest(http.MethodDelete, "/articles?id="+tt.id, nil)
			rec := httptest.NewRecorder()
			h.Article(rec, req)

			test.Eq(t, tt.want, rec.Code)
		})
	}
}