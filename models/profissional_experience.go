package models

import (
	"github.com/jinzhu/gorm"
)

type ProfessionalExperience struct {
	gorm.Model
	UserID      uint
	JobTitleID  uint
	BusinessID  uint
	StartDate   string
	EndDate     string
	Description string
}
