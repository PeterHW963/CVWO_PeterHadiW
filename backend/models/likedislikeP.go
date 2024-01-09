package models

import (
	"gorm.io/gorm"
)

type LikeDislikePost struct {
	gorm.Model
	PostId uint   `json:"postid"`
	UserId uint   `json:"userid"`
	Status string `json:"status"`
}
