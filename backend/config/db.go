package config

import (
	"github.com/PeterHW963/CVWO/backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:peterhw963@tcp(localhost:3306)/cvwo?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Post{})
	db.AutoMigrate(&models.Tag{})
	db.AutoMigrate(&models.Thread{})
	// db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.LikeDislikePost{})
	db.AutoMigrate(&models.LikeDislikeComment{})
	DB = db
	DB.Exec("ALTER TABLE posts ADD CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)")
}
