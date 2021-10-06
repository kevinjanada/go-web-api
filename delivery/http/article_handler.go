package http

import (
	"github.com/gin-gonic/gin"
	//"go-web-api/domain"
	"net/http"
)

type ArticleHandler struct {
	//TODO: Implement ArticleUsecase
	//AUsecase domain.ArticleUsecase
}

func NewArticleHandler(r *gin.Engine /*TODO: Implement ArticleUsecase, us domain.ArticleUsecase*/) {
	handler := &ArticleHandler{
		// TODO: Implement ArticleUsecase
		//AUsecase: us,
	}
	r.GET("/articles", handler.Fetch)
	r.POST("/articles", handler.Store)
	r.GET("/articles/:id", handler.GetByID)
	r.DELETE("/articles/:id", handler.Delete)
}

func (a *ArticleHandler) Fetch(c *gin.Context) {
	c.String(http.StatusOK, "TODO: Fetch")
}

func (a *ArticleHandler) Store(c *gin.Context) {
	c.String(http.StatusOK, "TODO: Store")
}

func (a *ArticleHandler) GetByID(c *gin.Context) {
	c.String(http.StatusOK, "TODO: GetByID")
}

func (a *ArticleHandler) Delete(c *gin.Context) {
	c.String(http.StatusOK, "TODO: Delete")
}
