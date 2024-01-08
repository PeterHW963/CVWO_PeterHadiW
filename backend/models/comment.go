package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	PostId  int    `json:"postid"`
	UserId  int    `json:"userid"`
	Content string `json:"content"`
}
