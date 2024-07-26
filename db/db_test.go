package db

import "testing"

func TestNew(t *testing.T) {
	c := &Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Database: "my_blog_test",
		SSL:      "disable",
	}
	db, err := New(c)
	if err != nil {
		t.Fatal("データベースの設定に失敗しました")
	}
	if db == nil {
		t.Fatal("dbがnilです")
	}
	_, err = db.Query("SELECT 1;")
	if err != nil {
		t.Fatal("Queryで失敗しました:", err)
	}
	if err := db.Close(); err != nil {
		t.Fatal("Closeで失敗しました:", err)
	}
}
