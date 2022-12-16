package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"library.demo/rent/database"
	"library.demo/rent/models"
)

func AddUserForm(c *gin.Context) {

	c.HTML(http.StatusOK, "add-user-form.html", gin.H{
		"title": "create new Client",
	})
}

func AddUserHandler(c *gin.Context) {

	client := models.Client{}

	c.Bind(&client)

	result := database.Instance.Create(&client)
	if result.Error != nil {
		c.HTML(http.StatusInternalServerError, "error.html", gin.H{
			"error": result.Error.Error(),
			"title": "error : 500",
		})
		return
	}

	c.Redirect(http.StatusFound, "/")

}
