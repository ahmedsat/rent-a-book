package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"library.demo/rent/database"
	"library.demo/rent/models"
)

func UpdateBookForm(c *gin.Context) {
	bookID, ok := c.Params.Get("id")
	if !ok {
		c.Redirect(http.StatusFound, "/")
		return
	}

	book := models.Book{}

	result := database.Instance.Find(&book, bookID)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}

	c.HTML(http.StatusOK, "update-book-form.html", gin.H{
		"title": "update new Book",
		"book":  book,
	})
}
func UpdateBookHandler(c *gin.Context) {

	book := models.Book{}

	database.Instance.Find(&book, c.Param("id"))

	err := c.Bind(&book)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err,
			"title": "error : 400",
		})
		return
	}

	result := database.Instance.Save(&book)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}

	c.Redirect(http.StatusFound, "/")

	// // c.JSON(http.StatusOK, gin.H{
	// // 	"url":    c.Request.URL.Path,
	// // 	"book": book,
	// // })
}
