package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"library.demo/rent/database"
	"library.demo/rent/models"
)

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	result := database.Instance.Delete(&models.BookItem{}, "book_refer = ?", id)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}
	result = database.Instance.Delete(&models.Book{}, id)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}
	c.Redirect(http.StatusFound, "/")
}
