package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"library.demo/rent/database"
	"library.demo/rent/models"
)

func Home(c *gin.Context) {

	books := []models.Book{}
	bookItems := []models.BookItem{}
	clients := []models.Client{}

	result := database.Instance.Find(&books)
	if result.Error != nil {
		log.Println(result.Error.Error())
	}
	booksCount := result.RowsAffected

	result = database.Instance.Find(&bookItems)
	if result.Error != nil {
		log.Println(result.Error.Error())
	}
	bookItemsCount := result.RowsAffected

	result = database.Instance.Model(models.Client{}).Preload("BookItems").Find(&clients)
	if result.Error != nil {
		log.Println(result.Error.Error())
	}
	clientsCount := result.RowsAffected

	// scratches
	// for _, client := range clients {

	// 	println(len(client.BookItems))
	// }
	// end scratches

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":          "book renter | main page",
		"booksCount":     booksCount,
		"bookItemsCount": bookItemsCount,
		"clientsCount":   clientsCount,
		"clients":        clients,
		"books":          books,
		"bookItems":      bookItems,
	})
}
