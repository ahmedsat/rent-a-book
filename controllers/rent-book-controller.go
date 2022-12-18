package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"library.demo/rent/database"
	"library.demo/rent/models"
)

func RentBook(c *gin.Context) {

	request := models.RentRequest{}

	err := c.Bind(&request)

	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err.Error(),
			"title": "error : 400",
		})
		return
	}

	bookItem := models.BookItem{}
	result := database.Instance.Find(&bookItem, request.Item)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}
	if bookItem.ID == 0 {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"error": "there is no book item with id:" + fmt.Sprint(request.Item),
			"title": "error : 500",
		})
		return
	}

	client := models.Client{}
	result = database.Instance.Find(&client, request.Client)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}
	if client.ID == 0 {
		c.HTML(http.StatusNotFound, "error.html", gin.H{
			"error": "there is no client with id:" + fmt.Sprint(request.Item),
			"title": "error : 500",
		})
		return
	}

	// how rent the book

	if !bookItem.RentedAt.Valid {
		bookItem.RentedAt = gorm.DeletedAt{Time: time.Now(), Valid: true}
		bookItem.RenterRefer = client.ID
		result = database.Instance.Save(&bookItem)
		if result.Error != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"error": result.Error.Error(),
				"title": "error : 500",
			})
			return
		}
		c.Redirect(http.StatusFound, "/client/"+fmt.Sprint(client.ID))
	}

	if bookItem.RenterRefer != client.ID {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "this book is already rented",
			"title": "error : 400",
		})
		return
	}

	// book is rented by user then return it

	bookItem.RentedAt = gorm.DeletedAt{Valid: false}
	bookItem.RenterRefer = 0
	result = database.Instance.Save(&bookItem)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}
	c.Redirect(http.StatusFound, "/client/"+fmt.Sprint(client.ID))
}
