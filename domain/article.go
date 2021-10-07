package domain

import (
	"context"
	"time"
)

type Article struct {
	ID      int64  `json:"id"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	//Author    Author    `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// ArticleUsecase represent the article's usecases
type ArticleUsecase interface {
	Fetch(ctx context.Context, num int64) (res []Article, err error)
	GetByID(ctx context.Context, id int64) (res Article, err error)
	Update(ctx context.Context, ar *Article) (err error)
	Store(ctx context.Context, ar *Article) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

// ArticleRepository represent the article's repository contract
type ArticleRepository interface {
	Fetch(ctx context.Context, num int64) (res []Article, err error)
	GetByID(ctx context.Context, id int64) (res Article, err error)
	Update(ctx context.Context, ar *Article) (err error)
	Store(ctx context.Context, ar *Article) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
