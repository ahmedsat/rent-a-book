package models

type RentRequest struct {
	Client uint `form:"client"`
	Item   uint `form:"item"`
}
