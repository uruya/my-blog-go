package article

import (
	"database/sql"
	"fmt"
	"time"
)

type Article struct {
	ID      int
	Title   string
	Content string
	Created time.Time
}

type Service struct {
	db *sql.DB
}

func New(db *sql.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) GetAll() ([]Article, error) {
	rows, err := s.db.Query("SELECT * FROM article ORDER BY id DESC LIMIT 100;")
	if err != nil {
		return nil, fmt.Errorf("クエリが失敗しました: %w", err)
	}
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var a Article
		if err := rows.Scan(&a.ID, &a.Title, &a.Content, &a.Created); err != nil {
			return nil, fmt.Errorf("スキャンに失敗しました: %w", err)
		}
		articles = append(articles, a)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return articles, nil
}

func (s *Service) Get(id int) (*Article, error) {
	row := s.db.QueryRow("SELECT * FROM article WHERE id = $1;", id)

	var a Article
	if err := row.Scan(&a.ID, &a.Title, &a.Content, &a.Created); err != nil {
		return nil, fmt.Errorf("スキャンに失敗しました: %w", err)
	}
	if err := row.Err(); err != nil {
		return nil, err
	}
	return &a, nil
}

func (s *Service) Create(title, content string) (int, error) {
	row := s.db.QueryRow(
		"INSERT INTO article (id, title, content) VALUES (nextval('article_id'), $1, $2) RETURNING id;",
		title, content,
	)

	var id int
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("インサートに失敗しました: %w", err)
	}

	if err := row.Err(); err != nil {
		return 0, fmt.Errorf("インサートに失敗しました: %w", err)
	}
	return id, nil
}

func (s *Service) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM article WHERE id=$1;", id)
	if err != nil {
		return fmt.Errorf("削除に失敗しました: %w", err)
	}
	return nil
}