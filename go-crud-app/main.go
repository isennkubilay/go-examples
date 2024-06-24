package main

import (
	"go-crud-app/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.ListBooks)
	r.GET("/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.html", nil)
	})
	r.POST("/create", handlers.CreateBook)
	r.GET("/update/:id", handlers.UpdateBookForm)
	r.POST("/update/:id", handlers.UpdateBook)
	r.GET("/delete/:id", handlers.DeleteBook)

	r.Run()
}
