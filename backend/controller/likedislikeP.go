package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/PeterHW963/CVWO/backend/config"
	"github.com/PeterHW963/CVWO/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt"
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
	type JWTToken struct {
		TokenString string `json:"stringToken"`
		PostId      uint   `json:"postid"`
		UserId      uint   `json:"userId"`
		Status      string `json:"status"`
	}

	var token JWTToken
	// c.ShouldBindJSON(&token)
	if err := c.ShouldBindBodyWith(&token, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// fmt.Print(token.TokenString)
	if token.TokenString == "" {
		c.String(200, "couldnt get cookie")
		return
	}

	result, err := jwt.Parse(token.TokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("KEY")), nil

	})

	if err != nil {
		c.String(200, "Token Parsing Failed")
		return
	}

	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		if float64(time.Now().Unix()) > claims["expires"].(float64) {
			c.String(200, "Token expired")
			return
		}
		var count int64
		var currentUser models.User
		config.DB.First(&currentUser, "id=?", claims["subject"]).Count(&count)

		if count == 0 {
			c.String(200, "User not found")
			c.Abort()
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
		c.Set("currentUser", currentUser)

	}

	upVote := models.LikeDislikePost{}

	var upVoteCount int64

	config.DB.Where("postid =? AND status = Up AND UserId=?", token.PostId, token.UserId).First(&upVote).Count(&upVoteCount)
	if upVoteCount == 0 {
		config.DB.Create(&models.LikeDislikePost{
			PostId: token.PostId,
			Status: token.Status,
			UserId: token.UserId,
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
	type JWTToken struct {
		TokenString string `json:"stringToken"`
		PostId      uint   `json:"postid"`
		UserId      uint   `json:"userId"`
		Status      string `json:"status"`
	}

	var token JWTToken
	// c.ShouldBindJSON(&token)
	if err := c.ShouldBindBodyWith(&token, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// fmt.Print(token.TokenString)
	if token.TokenString == "" {
		c.String(200, "couldnt get cookie")
		return
	}

	result, err := jwt.Parse(token.TokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("KEY")), nil

	})

	if err != nil {
		c.String(200, "Token Parsing Failed")
		return
	}

	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		if float64(time.Now().Unix()) > claims["expires"].(float64) {
			c.String(200, "Token expired")
			return
		}
		var count int64
		var currentUser models.User
		config.DB.First(&currentUser, "id=?", claims["subject"]).Count(&count)

		if count == 0 {
			c.String(200, "User not found")
			c.Abort()
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
		c.Set("currentUser", currentUser)

	}

	downVote := models.LikeDislikePost{}

	var downVoteCount int64

	config.DB.Where("postid =? AND status = Down AND UserId=?", token.PostId, token.UserId).First(&downVote).Count(&downVoteCount)
	if downVoteCount == 0 {
		config.DB.Create(&models.LikeDislikePost{
			PostId: token.PostId,
			Status: token.Status,
			UserId: token.UserId,
		})
		c.JSON(200, "Downvoted")
		return
	}

	if downVote.Status == "Down" {
		config.DB.Delete(&downVote)
		c.JSON(200, "Downvote Removed")
		return
	}

	downVote.Status = "Down"

	config.DB.Save(&downVote)
	c.JSON(200, "Downvoted")

}
