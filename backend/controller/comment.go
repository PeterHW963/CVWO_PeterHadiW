package controller

// create comment

// delete comment
// get comment
// edit comments

import (
	"github.com/PeterHW963/CVWO/backend/config"
	"github.com/PeterHW963/CVWO/backend/models"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	config.DB.Create(&comment)
	c.JSON(200, comment)
}

func UpdateComment(c *gin.Context) {
	var comment, newData models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
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

	if currentUser.ID != uint(comment.UserId) {
		c.JSON(403, gin.H{"error": "Forbidden"})
		c.Abort()
		return
	}

	config.DB.Where("id = ?", comment.ID).First(&newData)
	if newData.Content != comment.Content && newData.Content != "" {
		newData.Content = comment.Content

	}

	config.DB.Save(&newData)
	c.JSON(200, newData)
}

func DeleteComment(c *gin.Context) {
	var data struct {
		ID int `json:"id"`
	}

	c.ShouldBindJSON(&data)
	var comment models.Comment
	var count int64
	config.DB.Model(models.Comment{}).Where("id = ?", data.ID).First(&comment).Count(&count)
	if count > 0 {
		config.DB.Delete(&comment)
		c.JSON(200, comment)
		return
	}
	c.JSON(200, "comment not found")
}
