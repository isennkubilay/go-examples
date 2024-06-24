package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListBooks(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"books": books,
	})
}
