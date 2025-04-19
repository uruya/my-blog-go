# my-blog-go

Go + PostgreSQL で作ったシンプルなブログ API & サーバーです。  
現在は Railway でホスティング中 👉 **https://my-blog-go-production.up.railway.app**

## Features
- 記事の CRUD
- サーバーサイドレンダリング (html/template)
- Docker Compose でローカル DB 立ち上げ

## Tech Stack
| Layer | Tech |
|-------|------|
| Backend | Go 1.22 |
| DB | PostgreSQL 15 |
| Hosting | Railway (Free Tier) |

## Quick Start

```bash
git clone https://github.com/uruya/my-blog-go.git
cd my-blog-go

# 1. DB
docker compose up -d

# 2. App
go run main.go        # localhost:8080
