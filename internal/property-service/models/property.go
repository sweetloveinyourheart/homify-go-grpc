package models

import "gorm.io/gorm"

type Property struct {
	gorm.Model
	HostId      uint
	Title       string
	Description string
	Price       float32
	IsAvailable bool
	Destination Destination `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Category    []Category  `gorm:"many2many:property_categories;"`
	Amenity     []Amenity   `gorm:"many2many:property_amenities;"`
}
