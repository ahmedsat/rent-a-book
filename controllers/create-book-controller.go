package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddBookForm(c *gin.Context) {
	c.HTML(http.StatusOK, "add-book-form.html", gin.H{
		"title": "create new Book",
	})
}

func AddBookHandler(c *gin.Context) {}