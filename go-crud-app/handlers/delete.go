package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for i, b := range books {
		if b.ID == id {
			books = append(books[:i], books[i+1:]...)
			break
		}
	}

	c.Redirect(http.StatusFound, "/")
}
