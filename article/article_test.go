package article

import (
	"my-blog/test"
	"strconv"
	"testing"
)

func TestGetAll(t *testing.T) {
	testdb := test.DB(t)
	defer test.Close(t, testdb)

	s := New(testdb)
	got, err := s.GetAll()
	if err != nil {
		t.Fatal("記事の取得に失敗しました:", err)
	}

	if len(got) != 4 {
		t.Fatal("記事数が4ではありません:", len(got))
	}
	test.DB(t)
	test.Eq(t, "自己紹介", got[0].Title)
	test.Eq(t, "こんなことがありました", got[1].Title)
	test.Eq(t, "仕事について", got[2].Title)
	test.Eq(t, "ブログはじめました", got[3].Title)
}

func TestGet(t *testing.T) {
	testdb := test.DB(t)
	defer test.Close(t, testdb)

	s := New(testdb)
	for i := 1; i < 4; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got, err := s.Get(i)
			if err != nil {
				t.Fatal("記事の取得に失敗しました:", err)
			}
			if got == nil {
				t.Fatal("gotがnilです。")
			}
			if got.Title == "" {
				t.Fatal("タイトルが空です。")
			}
			if got.Content == "" {
				t.Fatal("中身が空です。")
			}
		})
	}

	t.Run("not found", func(t *testing.T) {
		got, err := s.Get(-1)
		if err == nil {
			t.Fatal("エラーがnilです。")
		}
		if got != nil {
			t.Fatal("gotがnilです。")
		}
	})
}

func TestCreateArticle(t *testing.T) {
	testdb := test.DB(t)
	defer test.Clear(t, testdb)
	s := New(testdb)

	id, err := s.Create("サンプルタイトル", "これはテスト用サンプルテキストです。")
	if err != nil {
		t.Fatal("failed to create article:", err)
	}

	// idのテスト
	test.Eq(t, 10000, id)

	// 作成したarticleを取得
	got, err := s.Get(id)
	if err != nil {
		t.Fatal("failed to get article:", err)
	}

	// Createに適した内容どおりか確認
	test.Eq(t, "サンプルタイトル", got.Title)
	test.Eq(t, "これはテスト用サンプルテキストです。", got.Content)
}

func TestDeleteArticle(t *testing.T) {
	testdb := test.DB(t)
	defer test.Close(t, testdb)
	s := New(testdb)

	err := s.Delete(1)
	if err != nil {
		t.Fatal("failed to create article:", err)
	}

	// articleを取得
	got, err := s.GetAll()
	if err != nil {
		t.Fatal("failed to get article:", err)
	}

	// 数が減っていればOK
	test.Eq(t, 3, len(got))

}