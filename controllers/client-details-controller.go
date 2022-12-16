package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"library.demo/rent/database"
	"library.demo/rent/models"
)

func ClientDetails(c *gin.Context) {
	clientID := c.Param("id")
	client := models.Client{}
	result := database.Instance.Model(client).Preload("BookItems").Find(&client, clientID)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}

	rentedBooksMap := make(map[uint]string)

	for _, bookItem := range client.BookItems {

		book := models.Book{}
		result = database.Instance.Find(&book, bookItem.BookRefer)
		if result.Error != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"error": result.Error.Error(),
				"title": "error : 500",
			})
			return
		}
		rentedBooksMap[bookItem.BookRefer] = book.Title
	}

	// c.HTML(http.StatusOK, "client-details.html", gin.H{
	// 	"title":  "book renter | client details",
	// 	"client": client,
	// 	"books":  rentedBooksMap,
	// })

	c.JSON(http.StatusOK, gin.H{
		"title":  "book renter | client details",
		"client": client,
		"books":  rentedBooksMap,
	})

}
