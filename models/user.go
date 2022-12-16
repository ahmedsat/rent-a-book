package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model

	Name    string `form:"name" gorm:"default:no name"`
	Email   string `form:"email" gorm:"unique"`
	Phone   string `form:"phone"`
	Address string `form:"address"`

	BookItems []BookItem `gorm:"foreignKey:RenterRefer;"`
}
