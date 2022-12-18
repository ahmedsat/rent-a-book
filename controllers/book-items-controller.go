package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"library.demo/rent/database"
	"library.demo/rent/models"
)

func AddBookItem(c *gin.Context) {
	id := 0

	_, err := fmt.Sscan(c.Param("id"), &id)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
			"title": "error : 400",
		})
		return
	}

	bookItem := models.BookItem{
		BookRefer: uint(id),
	}

	result := database.Instance.Create(&bookItem)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}

	c.Redirect(http.StatusFound, "/book/"+c.Param("id"))

}

func DeleteBookItem(c *gin.Context) {
	bookItem := models.BookItem{}
	result := database.Instance.Delete(&bookItem, c.Param("id"))
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}
	c.Redirect(http.StatusFound, "/")

}
