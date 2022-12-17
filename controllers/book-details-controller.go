package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"library.demo/rent/database"
	"library.demo/rent/models"
)

func BookDetails(c *gin.Context) {
	bookID := c.Param("id")
	book := models.Book{}
	result := database.Instance.Model(book).Preload("BookItems").Find(&book, bookID)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}

	bookItemsCount := len(book.BookItems)
	rentedBooksCount := 0
	for _, bookItem := range book.BookItems {
		if !bookItem.RentedAt.Valid {
			rentedBooksCount++
		}
	}

	c.HTML(http.StatusOK, "book-details.html", gin.H{
		"title":            "book renter | client details",
		"book":             book,
		"bookItemsCount":   bookItemsCount,
		"rentedBooksCount": rentedBooksCount,
	})

	// c.JSON(http.StatusOK, gin.H{
	// 	"title": "book renter | client details",
	// 	"book":  book,
	// })

}
