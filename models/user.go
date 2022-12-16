package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model

	Name    string `json:"name" gorm:"default:no name"`
	Email   string `json:"email" gorm:"unique"`
	Phone   string `json:"Phone"`
	Address string `json:"address"`
}
