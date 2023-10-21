package booksRouter

import (
	booksModule "main/GO_RESTAPI/modules/books"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(context *gin.Context) {
	books := booksModule.GetBooks()
	context.JSON(200, books)
}
