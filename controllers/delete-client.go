package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"library.demo/rent/database"
	"library.demo/rent/models"
)

func DeleteClient(c *gin.Context) {

	id := c.Param("id")

	updateItems := []models.BookItem{}

	result := database.Instance.Where("renter_refer = ?", id).Find(&updateItems)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}

	for _, item := range updateItems {
		item.RentedAt = gorm.DeletedAt{Valid: false}
		item.RenterRefer = 0
		result = database.Instance.Save(&item)
		if result.Error != nil {
			c.HTML(http.StatusInternalServerError, "error.html", gin.H{
				"error": result.Error.Error(),
				"title": "error : 500",
			})
			return
		}
	}

	result = database.Instance.Delete(&models.Client{}, c.Param("id"))
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}
	c.Redirect(http.StatusFound, "/")
}
