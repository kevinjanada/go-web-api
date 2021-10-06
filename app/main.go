package main

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"

	"github.com/gin-gonic/gin"
	"go-web-api/delivery/http"
)

func main() {
	dsn := "postgres://postgres:@localhost:5432/test?sslmode=disable"
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	//db := bun.NewDB(sqldb, pgdialect.New())
	_ = bun.NewDB(sqldb, pgdialect.New())

	r := gin.Default()
	/*
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	*/
	http.NewArticleHandler(r)
	r.Run()
}
