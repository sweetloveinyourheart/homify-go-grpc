package models

import "gorm.io/gorm"

type Listing struct {
	gorm.Model
	HostId      uint
	Title       string
	Description string
	Price       string
	IsAvailable bool
	Category    []Category `gorm:"many2many:listing_categories;"`
	Amenity     []Amenity  `gorm:"many2many:listing_amenities;"`
}
