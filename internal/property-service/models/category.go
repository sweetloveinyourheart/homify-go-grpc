package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"unique"`
	IsAvailable bool   `gorm:"default:true"`
	IconURL     string
}
