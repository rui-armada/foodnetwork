package models

import "github.com/jinzhu/gorm"

type Service struct {
	gorm.Model
	Name        string
	Description string
	BusinessID  uint
}
