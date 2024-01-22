package models

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	HostId      uint
	Title       string
	Description string
	Price       string
	IsAvailable bool
	Category    []Category `gorm:"many2many:property_categories;"`
	Amenity     []Amenity  `gorm:"many2many:property_amenities;"`
}
