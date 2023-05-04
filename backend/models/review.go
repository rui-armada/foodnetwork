package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	Text            string
	UserPublisherID uint
	UserID          uint
	ProductID       uint
	ServiceID       uint
	BusinessID      uint
}
