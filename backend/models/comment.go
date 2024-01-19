package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID      uint   `json:"id"`
	PostId  uint   `json:"postid"`
	UserId  uint   `json:"userid"`
	Content string `json:"content"`
}
