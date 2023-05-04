package models

import "gorm.io/gorm"

type Rating struct {
	gorm.Model
	Value      float64
	UserID     uint
	ProductID  uint
	ServiceID  uint
	BusinessID uint
}
