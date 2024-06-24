package handlers

import (
	"net/http"
	"strconv"

	"go-crud-app/models"

	"github.com/gin-gonic/gin"
)

func UpdateBookForm(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var book models.Book
	for _, b := range books {
		if b.ID == id {
			book = b
			break
		}
	}

	c.HTML(http.StatusOK, "update.html", gin.H{
		"book": book,
	})
}

func UpdateBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	title := c.PostForm("title")
	author := c.PostForm("author")

	for i, b := range books {
		if b.ID == id {
			books[i].Title = title
			books[i].Author = author
			break
		}
	}

	c.Redirect(http.StatusFound, "/")
}
