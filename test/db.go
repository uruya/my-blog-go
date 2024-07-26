package test

import (
	"database/sql"
	"my-blog/db"
	"testing"
)

func DB(t *testing.T) *sql.DB {
	t.Helper()
	c := &db.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Database: "my_blog_test",
		SSL:      "disable",
	}
	d, err := db.New(c)
	if err != nil {
		t.Fatal(err)
	}
	Clear(t, d)

	_, err = d.Exec(`
	INSERT INTO
		article (id, title, content)
	VALUES
		(1, 'ブログはじめました', 'このブログではわたしの個人的な事柄について書くつもりです。'),
		(2, '仕事について', 'わたしは新卒のころからずっと続けている仕事があります。'),
		(3, 'こんなことがありました', '先日、散歩をしていたときに変な出来事に遭遇しました。'),
		(4, '自己紹介', '今更ですが、わたしの自己紹介をさせてください。')
	`)
	if err != nil {
		t.Fatal(err)
	}
	return d
}

func Clear(t *testing.T, db *sql.DB) {
	t.Helper()
	if _, err := db.Exec("TRUNCATE article;"); err != nil {
		t.Fatal(err)
	}
	if _, err := db.Exec("SELECT SETVAL ('article_id', 10000, false);"); err != nil {
		t.Fatal(err)
	}
}

func Close(t *testing.T, db *sql.DB) {
	t.Helper()
	Clear(t, db)
	if err := db.Close(); err != nil {
		t.Fatal(err)
	}
}
