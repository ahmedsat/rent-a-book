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
	RentedAt    gorm.DeletedAt
	BookRefer   uint
	RenterRefer uint
	Book        Book   `gorm:"foreignKey:BookRefer;"`
	Renter      Client `gorm:"foreignKey:RenterRefer;"`
}
