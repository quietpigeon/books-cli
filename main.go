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
		v1.GET("/")
	}
	r.Run()
}
