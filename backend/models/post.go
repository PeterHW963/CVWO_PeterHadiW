package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	// PostId   int    `json:"postid"`
	ID       uint   `json:"id" gorm:"primaryKey"`
	UserId   uint   `json:"userid"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	TagId    *uint  `json:"tagid"`
	ThreadId *uint  `json:"threadid"`
}
