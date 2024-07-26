package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSL      string
}

func (c *Config) String() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.Database,
		c.SSL,
	)
}

func New(config *Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.String())
	if err != nil {
		return nil, fmt.Errorf("データベースの接続に失敗しました: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("pingに失敗しました: %w", err)
	}

	return db, nil
}
