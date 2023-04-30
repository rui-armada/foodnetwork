package models

import "gorm.io/gorm"

type BusinessUser struct {
	gorm.Model
	UserID     uint
	BusinessID uint
	JobTitleID uint
}
