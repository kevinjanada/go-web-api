package usecase

import (
	"context"
	"go-web-api/domain"
)

type articleUsecase struct {
	articleRepo domain.ArticleRepository
}

func NewArticleUsecase(articleRepo domain.ArticleRepository) domain.ArticleUsecase {
	return &articleUsecase{
		articleRepo,
	}
}

func (a *articleUsecase) Fetch(ctx context.Context, num int64) (res []domain.Article, err error) {
	return
}

func (a *articleUsecase) GetByID(ctx context.Context, id int64) (res domain.Article, err error) {
	return
}

func (a *articleUsecase) Update(ctx context.Context, article *domain.Article) (err error) {
	return
}

func (a *articleUsecase) Store(ctx context.Context, article *domain.Article) (err error) {
	return
}

func (a *articleUsecase) Delete(ctx context.Context, id int64) (err error) {
	return
}
