package controller

// create comment

// delete comment
// get comment
// edit comments

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

func CreateComment(c *gin.Context) {
	type JWTToken struct {
		TokenString string `json:"stringToken"`
		PostId      uint   `json:"id"`
		UserId      uint   `json:"userId"`
		Content     string `json:"content"`
	}

	var token JWTToken

	if err := c.ShouldBindBodyWith(&token, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

		var comment models.Comment

		comment.PostId = token.PostId
		comment.UserId = token.UserId
		comment.Content = token.Content
		config.DB.Create(&comment)
		c.JSON(200, comment)
	}

}

func UpdateComment(c *gin.Context) {

	type JWTToken struct {
		TokenString string `json:"stringToken"`
		ID          uint   `json:"id"`
		Content     string `json:"content"`
	}

	var token JWTToken

	if err := c.ShouldBindBodyWith(&token, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

		var comment, newData models.Comment

		comment.ID = token.ID
		comment.Content = token.Content

		config.DB.Where("id = ?", comment.ID).First(&newData)

		if currentUser.ID != newData.UserId {
			c.JSON(403, gin.H{"error": "Forbidden: You can only update your own posts"})
			return
		}

		newData.Content = comment.Content
		config.DB.Save(&newData)
		c.JSON(200, newData)
	}
}

func DeleteComment(c *gin.Context) {

	type JWTToken struct {
		TokenString string `json:"stringToken"`
		ID          uint   `json:"id"`
	}

	var token JWTToken
	if err := c.ShouldBindBodyWith(&token, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	var currentUser models.User
	if claims, ok := result.Claims.(jwt.MapClaims); ok && result.Valid {
		if float64(time.Now().Unix()) > claims["expires"].(float64) {
			c.String(200, "Token expired")
			return
		}
		var count int64
		config.DB.First(&currentUser, "id=?", claims["subject"]).Count(&count)

		if count == 0 {
			c.String(200, "User not found")
			c.Abort()
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}
	}

	if currentUser.Role == "Admin" {

		var comment models.Comment
		comment.ID = token.ID
		var count int64
		config.DB.Model(models.Comment{}).Where("id = ?", token.ID).First(&comment).Count(&count)
		if count > 0 {
			config.DB.Delete(&comment)
			c.JSON(200, comment)
			return
		}
		c.JSON(200, "comment not found")
	}

	c.JSON(403, gin.H{"error": "Forbidden"})
	c.Abort()
	c.Redirect(http.StatusTemporaryRedirect, "/")

}
