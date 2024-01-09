package models

import (
	"gorm.io/gorm"
)

type Thread struct {
	gorm.Model
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
