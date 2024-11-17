package main

import (
	"books-cli/api"

	"github.com/gin-gonic/gin"
)

func main() {
	api.LoadBooks()
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/books", api.HandleGetBooks)
		v1.POST("/books", api.HandleAddBooks)
		v1.GET("/books/:id", api.HandleGetBooksByID)
		v1.PUT("/books/:id", api.HandleUpdateBookByID)
		v1.DELETE("/books/:id")
	}
	r.Run(":8080")
}
