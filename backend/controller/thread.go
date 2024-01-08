package controller

import (
	"fmt"

	"github.com/PeterHW963/CVWO/backend/config"
	"github.com/PeterHW963/CVWO/backend/models"
	"github.com/gin-gonic/gin"
)

func CreateThread(c *gin.Context) {
	var thread models.Thread
	if err := c.ShouldBindJSON(&thread); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	config.DB.Create(&thread)
	c.JSON(200, thread)
}

func GetThread(c *gin.Context) {
	var data struct {
		Search string `json:"search"`
	}

	c.ShouldBindJSON(&data)
	threads := []models.Thread{}
	config.DB.Where("name LIKE ?", "%"+data.Search+"%").Find(&threads)
	fmt.Println(threads)
	c.JSON(200, threads)
}

func UpdateThread(c *gin.Context) {
	var thread, newData models.Thread
	if err := c.ShouldBindJSON(&thread); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Where("id = ?", thread.ID).First(&newData)
	newData.Name = thread.Name
	newData.Description = thread.Description

	config.DB.Save(&newData)
	c.JSON(200, newData)
}

func DeleteThread(c *gin.Context) {
	var data struct {
		ID int `json:"id"`
	}

	c.ShouldBindJSON(&data)
	var thread models.Thread
	var count int64
	config.DB.Model(models.Thread{}).Where("id = ?", data.ID).First(&thread).Count(&count)
	if count > 0 {
		config.DB.Delete(&thread)
		c.JSON(200, thread)
		return
	}
	c.JSON(200, "thread not found")
}
