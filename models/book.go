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

	BookItems []BookItem `gorm:"foreignKey:BookRefer"`
}

type BookItem struct {
	gorm.Model
	RentedAt    gorm.DeletedAt
	BookRefer   uint
	RenterRefer uint
}
