package models

import "gorm.io/gorm"

type JobPost struct {
	gorm.Model
	Title       string
	Description string
	BusinessID  uint
}
