package models

import "gorm.io/gorm"

type Business struct {
	gorm.Model
	Name           string
	Address        string
	BusinessTypeID uint
}
