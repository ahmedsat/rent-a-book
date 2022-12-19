package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload" // load .env file
	"library.demo/rent/database"
	"library.demo/rent/routes"
)

func main() {

	
	AutoMigrate := os.Getenv("AUTO_MIGRATE")
	Seed := os.Getenv("SEED")
	// port := "localhost:" + os.Getenv("PORT")

	database.Conn()

	
	if AutoMigrate == "TRUE" {
		database.MigrateAll()
	}
	
	if Seed == "TRUE" {
		go database.Seed()
	}
	
	r := gin.Default()
	
	r.LoadHTMLGlob("templates/*")

	routes.RootRoutes(r)
	r.Run()

}
