package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload" // load .env file
	"library.demo/rent/database"
)

// Book ...
type Book struct {
	Title  string
	Author string
}

func main() {

	AutoMigrate := os.Getenv("AUTO_MIGRATE")
	port := "localhost:" + os.Getenv("PORT")

	database.Conn()

	r := gin.Default()
	r.LoadHTMLGlob("templates/**/*")

	if AutoMigrate == "TRUE" {
		database.MigrateAll()
	}

	// routes.RootRoutes(r)
	r.Run(port)

}
