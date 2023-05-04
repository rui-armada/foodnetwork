package models

import "gorm.io/gorm"

type BusinessType struct {
	gorm.Model
	Name string
}
