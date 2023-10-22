package booksRouter

import (
	booksModule "main/GO_RESTAPI/modules/books"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(context *gin.Context) {
	books := booksModule.GetBooks()
	context.JSON(200, books)
}

func GetBook(context *gin.Context) {
	id := context.Param("id")
	book := booksModule.GetBook(id)
	context.JSON(200, book)
}
