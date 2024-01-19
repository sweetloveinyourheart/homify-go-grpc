package models

import "gorm.io/gorm"

type Destination struct {
	gorm.Model
	Country   string
	City      string
	Latitude  string
	Longitude string
	ListingId uint
	Listing   Listing `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
