package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"library.demo/rent/database"
	"library.demo/rent/models"
)

func DeleteClient(c *gin.Context) {

	id := c.Param("id")

	client:=models.Client{}
	result:=database.Instance.Model(client).Preload("BookItems").Find(&client,id)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}

	if len(client.BookItems)>0{
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": "client has a rented books",
			"title": "error : 400",
		})
		return
	}


	result = database.Instance.Delete(&client, c.Param("id"))
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}
	c.Redirect(http.StatusFound, "/")
}
