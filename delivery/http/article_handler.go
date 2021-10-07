package http

import (
	"github.com/gin-gonic/gin"
	"go-web-api/domain"
	"net/http"
)

type ArticleHandler struct {
	AUsecase domain.ArticleUsecase
}

func NewArticleHandler(r *gin.Engine, us domain.ArticleUsecase) {
	handler := &ArticleHandler{
		AUsecase: us,
	}
	r.GET("/articles", handler.Fetch)
	r.POST("/articles", handler.Store)
	r.GET("/articles/:id", handler.GetByID)
	r.DELETE("/articles/:id", handler.Delete)
}

func (a *ArticleHandler) Fetch(ctx *gin.Context) {
	ctx.String(http.StatusOK, "TODO: Fetch")
}

func (a *ArticleHandler) Store(ctx *gin.Context) {
	ctx.String(http.StatusOK, "TODO: Store")
}

func (a *ArticleHandler) GetByID(ctx *gin.Context) {
	ctx.String(http.StatusOK, "TODO: GetByID")
}

func (a *ArticleHandler) Delete(ctx *gin.Context) {
	ctx.String(http.StatusOK, "TODO: Delete")
}
