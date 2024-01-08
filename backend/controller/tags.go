package controller

import (
	"fmt"

	"github.com/PeterHW963/CVWO/backend/config"
	"github.com/PeterHW963/CVWO/backend/models"
	"github.com/gin-gonic/gin"
)

func CreateTag(c *gin.Context) {
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	config.DB.Create(&tag)
	c.JSON(200, tag)
}

func GetTag(c *gin.Context) {
	var data struct {
		Search string `json:"search"`
	}

	c.ShouldBindJSON(&data)
	tags := []models.Tag{}
	config.DB.Where("name LIKE ?", "%"+data.Search+"%").Find(&tags)
	fmt.Println(tags)
	c.JSON(200, tags)
}

func UpdateTag(c *gin.Context) {
	var tag, newData models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Where("id = ?", tag.ID).First(&newData)
	newData.TagName = tag.TagName

	config.DB.Save(&newData)
	c.JSON(200, newData)
}

func DeleteTag(c *gin.Context) {
	var data struct {
		ID int `json:"id"`
	}

	c.ShouldBindJSON(&data)
	var tag models.Tag
	var count int64
	config.DB.Model(models.Tag{}).Where("id = ?", data.ID).First(&tag).Count(&count)
	if count > 0 {
		config.DB.Delete(&tag)
		c.JSON(200, tag)
		return
	}
	c.JSON(200, "tag not found")
}
