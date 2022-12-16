package routes

import (
	"github.com/gin-gonic/gin"
	"library.demo/rent/controllers"
)

func RootRoutes(r *gin.Engine) {

	r.Static("/assets", "./assets")
	// r.StaticFS("/more_static", http.Dir("my_file_system"))
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	r.GET("/", controllers.Home)

	r.GET("/client/add", controllers.AddUserForm)
	r.POST("/client/add", controllers.AddUserHandler)

	r.GET("/client/update", controllers.UpdateUserForm)
	r.GET("/client/update/:id", controllers.UpdateUserForm)
	r.POST("/client/update", controllers.UpdateUserHandler)
	r.POST("/client/update/:id", controllers.UpdateUserHandler)

	r.GET("/client/delete/:id", controllers.DeleteClient)

	r.GET("/client/:id", controllers.ClientDetails)

	r.GET("/test", mocController)

}
