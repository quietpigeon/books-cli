package main

import (
	"books-cli/api"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func setupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/books", func(ctx *gin.Context) { api.HandleGetBooks(ctx, db) })
		v1.POST("/books", func(ctx *gin.Context) { api.HandleAddBooks(ctx, db) })
		v1.GET("/books/:id", func(ctx *gin.Context) { api.HandleGetBooksByID(ctx, db) })
		v1.PUT("/books/:id", func(ctx *gin.Context) { api.HandleUpdateBookByID(ctx, db) })
		v1.DELETE("/books/:id", func(ctx *gin.Context) { api.HandleDeleteBookByID(ctx, db) })
	}
	return r
}

func main() {
	db, err := api.InitializeDB()
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}
	defer db.Close()

	r := setupRouter(db)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
