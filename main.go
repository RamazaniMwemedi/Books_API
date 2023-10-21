package main

import (
	booksRouter "main/GO_RESTAPI/controllers/books"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// Get all books
	router.GET("/books", func(ctx *gin.Context) {
		booksRouter.GetAllBooks(ctx)
	})
	// Get a
	
	router.Run(":8080")

}
