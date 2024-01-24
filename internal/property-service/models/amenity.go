package models

import "gorm.io/gorm"

type Amenity struct {
	gorm.Model
	Name        string `gorm:"unique"`
	IsAvailable bool   `gorm:"default:true"`
	IconURL     string
}
