package handlers

import (
	"net/http"

	"go-crud-app/models"

	"github.com/gin-gonic/gin"
)

var books []models.Book
var nextID int = 1

func CreateBook(c *gin.Context) {
	title := c.PostForm("title")
	author := c.PostForm("author")

	book := models.Book{
		ID:     nextID,
		Title:  title,
		Author: author,
	}
	nextID++

	books = append(books, book)
	c.Redirect(http.StatusFound, "/")
}
