package controller

import (
	"github.com/PeterHW963/CVWO/backend/config"
	"github.com/PeterHW963/CVWO/backend/models"
	"github.com/gin-gonic/gin"
)

func GetTotalVotesPost(c *gin.Context) {
	upVote := []models.LikeDislikePost{}
	var data struct {
		PostId int `json:"postid"`
	}
	c.ShouldBindJSON(&data)
	var countUp int64
	config.DB.Where("postid = ? AND status = Up", data.PostId).Find(&upVote).Count(&countUp)

	downVote := []models.LikeDislikePost{}
	var countDown int64
	config.DB.Where("postid = ? AND status = Down", data.PostId).Find(&downVote).Count(&countDown)

	TotalCount := countUp - countDown
	c.JSON(200, TotalCount)
}

func AddUpvotePost(c *gin.Context) {
	upVote := models.LikeDislikePost{}
	var data struct {
		PostId int    `json:"postid"`
		UserId int    `json:"userid"`
		Status string `json:"status"`
	}
	c.ShouldBindJSON(&data)
	var upVoteCount int64

	config.DB.Where("postid =? AND status = Up AND UserId=?", data.PostId, data.UserId).First(&upVote).Count(&upVoteCount)
	if upVoteCount == 0 {
		config.DB.Create(&models.LikeDislikePost{
			PostId: data.PostId,
			Status: data.Status,
			UserId: data.UserId,
		})
		c.JSON(200, "Upvoted")
		return
	}

	if upVote.Status == "Up" {
		config.DB.Delete(&upVote)
		c.JSON(200, "Upvote Removed")
		return
	}

	upVote.Status = "Up"

	config.DB.Save(&upVote)
	c.JSON(200, "Upvoted")

}

func AddDownvotePost(c *gin.Context) {
	downVote := models.LikeDislikePost{}
	var data struct {
		PostId int    `json:"postid"`
		UserId int    `json:"userid"`
		Status string `json:"status"`
	}
	c.ShouldBindJSON(&data)
	var downVoteCount int64

	config.DB.Where("postid =? AND status = Down AND UserId=?", data.PostId, data.UserId).First(&downVote).Count(&downVoteCount)
	if downVoteCount == 0 {
		config.DB.Create(&models.LikeDislikePost{
			PostId: data.PostId,
			Status: data.Status,
			UserId: data.UserId,
		})
		c.JSON(200, "Downvoted")
		return
	}

	if downVote.Status == "Up" {
		config.DB.Delete(&downVote)
		c.JSON(200, "Downvote Removed")
		return
	}

	downVote.Status = "Up"

	config.DB.Save(&downVote)
	c.JSON(200, "Downvoted")
}
