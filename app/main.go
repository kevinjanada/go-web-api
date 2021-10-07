package main

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/gin-gonic/gin"
	"go-web-api/delivery/http"
	"go-web-api/repository/psql"
	"go-web-api/usecase"
)

func main() {
	dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	db := bun.NewDB(sqldb, pgdialect.New())

	r := gin.Default()
	/*
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	*/
	articleRepo := psql.NewArticleRepository(db)
	articleUsecase := usecase.NewArticleUsecase(articleRepo)
	http.NewArticleHandler(r, articleUsecase)
	r.Run()
}
