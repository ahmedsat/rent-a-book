package routes

import (
	"github.com/gin-gonic/gin"
	"library.demo/rent/controllers"
)

func RootRoutes(r *gin.Engine) {

	r.Static("/assets", "./assets") // serve statics (static html, css, javascript, images, ...)
	// r.StaticFS("/more_static", http.Dir("my_file_system"))
	r.StaticFile("/favicon.ico", "./resources/favicon.ico") // serve favicon

	r.GET("/", controllers.Home) // root path - main page

	// clint routs
	r.GET("/client/:id", controllers.ClientDetails)

	r.GET("/client/add", controllers.AddClientForm)
	r.POST("/client/add", controllers.AddClientHandler)

	r.GET("/client/update/:id", controllers.UpdateClientForm)
	r.POST("/client/update/:id", controllers.UpdateClientHandler)

	r.GET("/client/delete/:id", controllers.DeleteClient)

	// book routs
	r.GET("/book/:id", controllers.BookDetails)

	r.GET("/book/add", controllers.AddBookForm)
	r.POST("/book/add", controllers.AddBookHandler)

	r.GET("/book/update/:id", controllers.UpdateBookForm)
	r.POST("/book/update/:id", controllers.UpdateBookHandler)

	r.GET("/book/delete/:id", controllers.DeleteBook)

	// book routs
	r.GET("/bookItem/add/:id", controllers.AddBookItem)
	r.GET("/bookItem/delete/:id", controllers.DeleteBookItem)

	// renting routs
	r.POST("/rent", controllers.RentBook)

	r.GET("/test", mocController)

}
