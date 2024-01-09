package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostId  uint   `json:"postid"`
	UserId  uint   `json:"userid"`
	Content string `json:"content"`
}
