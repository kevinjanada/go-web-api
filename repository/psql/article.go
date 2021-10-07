package psql

import (
	"context"
	"github.com/uptrace/bun"
	"go-web-api/domain"
)

type articleRepository struct {
	Conn *bun.DB
}

func NewArticleRepository(db *bun.DB) *articleRepository {
	return &articleRepository{
		Conn: db,
	}
}

func (a *articleRepository) Fetch(ctx context.Context, num int64) (res []domain.Article, err error) {
	return
}

func (a *articleRepository) GetByID(ctx context.Context, id int64) (res domain.Article, err error) {
	return
}

func (a *articleRepository) Update(ctx context.Context, ar *domain.Article) (err error) {
	return
}

func (a *articleRepository) Store(ctx context.Context, ar *domain.Article) (err error) {
	return
}

func (a *articleRepository) Delete(ctx context.Context, id int64) (err error) {
	return
}
