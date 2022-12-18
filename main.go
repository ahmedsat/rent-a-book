package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload" // load .env file
	"library.demo/rent/database"
	"library.demo/rent/routes"
)

func main() {

	// Generate()

	AutoMigrate := os.Getenv("AUTO_MIGRATE")
	Seed := os.Getenv("SEED")
	port := "localhost:" + os.Getenv("PORT")

	database.Conn()

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	// r.LoadHTMLFiles("Assets")

	if AutoMigrate == "TRUE" {
		database.MigrateAll()
	}

	if Seed == "TRUE" {
		go database.Seed()
	}

	routes.RootRoutes(r)
	r.Run(port)

}

// func loadTemplate() (*template.Template, error) {
// 	t := template.New("")
// 	for name, file := range Assets.Files {
// 		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
// 			continue
// 		}
// 		h, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			return nil, err
// 		}
// 		t, err = t.New(name).Parse(string(h))
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return t, nil
// }

// func Generate() {
// 	log.Println("start")
// 	g := assets.Generator{PackageName: "tmp", VariableName: "var"}

// 	if err := g.Add("templates"); err != nil {
// 		panic(err)
// 	}

// 	// This will write a go file to standard out. The generated go file
// 	// will reside in the g.PackageName package and will contain a
// 	// single variable g.VariableName of type assets.FileSystem containing
// 	// the whole file system.
// 	g.Write(os.Stdout)
// 	log.Println("end")
// }
