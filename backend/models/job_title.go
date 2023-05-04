package models

import "gorm.io/gorm"

type JobTitle struct {
	gorm.Model
	Name string
}

