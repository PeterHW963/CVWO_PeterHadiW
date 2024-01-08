package models

import (
	"gorm.io/gorm"
)

type LikeDislikePost struct {
	gorm.Model
	PostId int    `json:"postid"`
	UserId int    `json:"userid"`
	Status string `json:"status"`
}
