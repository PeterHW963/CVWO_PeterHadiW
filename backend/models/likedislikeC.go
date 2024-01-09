package models

import (
	"gorm.io/gorm"
)

type LikeDislikeComment struct {
	gorm.Model
	CommentId uint   `json:"commentid"`
	UserId    uint   `json:"userid"`
	Status    string `json:"status"`
}
