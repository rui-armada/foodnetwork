package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID      uint
	BusinessID  uint
	Description string
}
