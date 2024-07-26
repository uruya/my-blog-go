package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_Index(t *testing.T) {
	h, close := newHandler(t)
	defer close()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	h.Index(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatal("ステータスコードが200ではありません。", rec.Code)
	}

	t.Log(rec.Body)
	// トップページのタイトルのテスト
	if !strings.Contains(rec.Body.String(), `<h1 class="fs-2">わたしのブログ</h1>`) {
		t.Fatal("タイトルがありません。")
	}

	// 記事一覧のテスト
	titles := []string{"自己紹介", "こんなことがありました", "仕事について", "ブログはじめました"}
	for _, title := range titles {
		if !strings.Contains(rec.Body.String(), fmt.Sprintf("<h2>%s</h2>", title)) {
			t.Fatal("記事のタイトルがありません:", title)
		}
	}
}
