package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Year      int
	Title     string
	Writer    string
	Publisher string
	Location  string
	ImageURI  string
}

type BookItem struct {
	gorm.Model
	RentedAt gorm.DeletedAt
	Book     Book   `gorm:"foreignKey:ID;"`
	Renter   Client `gorm:"foreignKey:ID;"`
}
