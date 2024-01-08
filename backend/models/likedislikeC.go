package models

import (
	"gorm.io/gorm"
)

type LikeDislikeComment struct {
	gorm.Model
	CommentId int    `json:"commentid"`
	UserId    int    `json:"userid"`
	Status    string `json:"status"`
}
