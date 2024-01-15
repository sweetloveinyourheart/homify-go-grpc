package models

import (
	"database/sql"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName  string
	Gender    sql.NullString
	Birthday  sql.NullString
	Phone     sql.NullString
	AccountID uint
	Account   Account `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
