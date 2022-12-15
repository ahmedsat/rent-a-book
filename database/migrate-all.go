package database

func MigrateAll() {

	Instance.AutoMigrate(
	// &models.Image{},
	// &models.Category{},
	// &models.User{},
	// &models.Product{},
	// &models.CartItem{},
	// &models.Order{},
	)

}
