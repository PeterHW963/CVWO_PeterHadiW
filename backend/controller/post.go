package controller

// get post
// create post
// delete post
// udpate post

import (
	"fmt"

	"github.com/PeterHW963/CVWO/backend/config"
	"github.com/PeterHW963/CVWO/backend/models"
	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	config.DB.Create(&post)
	c.JSON(200, post)
}

func GetPost(c *gin.Context) {
	var data struct {
		Search string `json:"search"`
	}

	c.ShouldBindJSON(&data)
	posts := []models.Post{}
	config.DB.Where("name LIKE ?", "%"+data.Search+"%").Find(&posts)
	fmt.Println(posts)
	c.JSON(200, posts)
}

func UpdatePost(c *gin.Context) {
	var post, newData models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	currentUserInterface, exists := c.Get("currentUser")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	currentUser, ok := currentUserInterface.(models.User)
	if !ok {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		c.Abort()
		return
	}

	if currentUser.ID != post.UserId {
		c.JSON(403, gin.H{"error": "Forbidden"})
		c.Abort()
		return
	}

	config.DB.Where("id = ?", post.ID).First(&newData)
	if newData.Content != post.Content && newData.Content != "" {
		newData.Content = post.Content

	}

	config.DB.Save(&newData)
	c.JSON(200, newData)
}

func DeletePost(c *gin.Context) {
	var data struct {
		ID int `json:"id"`
	}

	c.ShouldBindJSON(&data)
	var post models.Post
	var count int64
	config.DB.Model(models.Post{}).Where("id = ?", data.ID).First(&post).Count(&count)
	if count > 0 {
		config.DB.Delete(&post)
		c.JSON(200, post)
		return
	}
	c.JSON(200, "post not found")
}
