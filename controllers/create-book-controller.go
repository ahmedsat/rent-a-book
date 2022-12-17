package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"library.demo/rent/database"
	"library.demo/rent/models"
)

func AddBookForm(c *gin.Context) {
	c.HTML(http.StatusOK, "add-book-form.html", gin.H{
		"title": "create new Book",
	})
}

func AddBookHandler(c *gin.Context) {

	book := models.Book{}

	err := c.Bind(&book)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err,
			"title": "error : 400",
		})
		return
	}

	result := database.Instance.Create(&book)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}

	c.Redirect(http.StatusFound, "/")

}
