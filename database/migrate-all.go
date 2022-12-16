package database

import "library.demo/rent/models"

func MigrateAll() {

	Instance.AutoMigrate(
		&models.BookItem{},
		&models.Client{},
		&models.Book{},
	)

}
