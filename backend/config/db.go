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
	DB.Exec("ALTER TABLE posts ADD CONSTRAINT fk_user_posts FOREIGN KEY (user_id) REFERENCES users(id)")
	DB.Exec("ALTER TABLE posts ADD CONSTRAINT fk_tag_posts FOREIGN KEY (tag_id) REFERENCES tags(id)")
	DB.Exec("ALTER TABLE posts ADD CONSTRAINT fk_thread_posts FOREIGN KEY (thread_id) REFERENCES threads(id)")
	DB.Exec("ALTER TABLE like_dislike_comments ADD CONSTRAINT fk_comment_LDcomments FOREIGN KEY (comment_id) REFERENCES comments(id)")
	DB.Exec("ALTER TABLE like_dislike_comments ADD CONSTRAINT fk_user_LDcomments FOREIGN KEY (user_id) REFERENCES users(id)")
	DB.Exec("ALTER TABLE like_dislike_posts ADD CONSTRAINT fk_user_LDposts FOREIGN KEY (user_id) REFERENCES users(id)")
	DB.Exec("ALTER TABLE like_dislike_posts ADD CONSTRAINT fk_post_LDposts FOREIGN KEY (post_id) REFERENCES posts(id)")
	DB.Exec("ALTER TABLE comments ADD CONSTRAINT fk_post_comments FOREIGN KEY (post_id) REFERENCES posts(id)")
	DB.Exec("ALTER TABLE comments ADD CONSTRAINT fk_user_comments FOREIGN KEY (user_id) REFERENCES users(id)")

}
