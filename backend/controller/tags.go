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

func CreateTag(c *gin.Context) {
	type JWTToken struct {
		TokenString string `json:"stringToken"`
		ID          uint   `json:"id"`
		TagName     string `json:"tagname"`
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

		var tag models.Tag

		tag.ID = token.ID
		tag.TagName = token.TagName
		config.DB.Create(&tag)
		c.JSON(200, tag)
	}

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

	type JWTToken struct {
		TokenString string `json:"stringToken"`
		ID          uint   `json:"id"`
		TagName     string `json:"tagname"`
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

		var tag, newData models.Tag

		tag.ID = token.ID
		tag.TagName = token.TagName
		config.DB.Where("id = ?", tag.ID).First(&newData)
		newData.TagName = tag.TagName

		config.DB.Save(&newData)
		c.JSON(200, newData)
	}

}

func DeleteTag(c *gin.Context) {

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

		var tag models.Tag
		tag.ID = token.ID
		var count int64
		config.DB.Model(models.Tag{}).Where("id = ?", token.ID).First(&tag).Count(&count)
		if count > 0 {
			config.DB.Delete(&tag)
			c.JSON(200, tag)
			return
		}
		c.JSON(200, "tag not found")
	}

	c.JSON(403, gin.H{"error": "Forbidden"})
	c.Abort()
	c.Redirect(http.StatusTemporaryRedirect, "/")

}
