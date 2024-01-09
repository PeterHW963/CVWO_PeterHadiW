package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	ID      uint   `json:"id" gorm:"primaryKey"`
	TagName string `json:"tagname"`
}
