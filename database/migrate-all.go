package database

import "library.demo/rent/models"

func MigrateAll() {

	Instance.AutoMigrate(
		&models.Book{},
		&models.BookItem{},
		&models.Client{},
	)

}
