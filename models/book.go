package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model

	Title     string `form:"name"`
	Writer    string `form:"writer"`
	Publisher string `form:"publisher"`
	Year      int    `form:"year"`
	Location  string `form:"location"`
	ImageURI  string `form:"image" gorm:"default:https://c1.wallpaperflare.com/preview/304/436/847/book-embossing-leather-empty.jpg"`

	BookItems []BookItem `gorm:"foreignKey:BookRefer"`
}

type BookItem struct {
	gorm.Model
	RentedAt    gorm.DeletedAt
	BookRefer   uint
	RenterRefer uint
}
