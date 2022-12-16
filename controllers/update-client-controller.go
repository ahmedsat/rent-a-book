package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"library.demo/rent/database"
	"library.demo/rent/models"
)

func UpdateUserForm(c *gin.Context) {
	clientID, ok := c.Params.Get("id")
	if !ok {
		c.Redirect(http.StatusFound, "/")
		return
	}

	client := models.Client{}

	result := database.Instance.Find(&client, clientID)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}

	c.HTML(http.StatusOK, "update-user-form.html", gin.H{
		"title":  "update new Client",
		"client": client,
	})
}
func UpdateUserHandler(c *gin.Context) {

	client := models.Client{}

	database.Instance.Find(&client, c.Param("id"))

	err := c.Bind(&client)
	if err != nil {
		c.HTML(http.StatusBadRequest, "error.html", gin.H{
			"error": err,
			"title": "error : 400",
		})
		return
	}

	result := database.Instance.Save(&client)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}

	c.Redirect(http.StatusFound, "/")

	// c.JSON(http.StatusOK, gin.H{
	// 	"url":    c.Request.URL.Path,
	// 	"client": client,
	// })
}
