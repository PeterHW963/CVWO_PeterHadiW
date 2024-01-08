package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	// PostId   int    `json:"postid"`
	UserId  uint   `json:"userid" gorm:"primary_key References:users(ID)"`
	Title   string `json:"title"`
	Content string `json:"content"`
	TagId   int    `json:"tagid"`
	TopicId int    `json:"topicid"`
}
