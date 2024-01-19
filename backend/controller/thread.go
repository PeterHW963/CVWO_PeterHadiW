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

func CreateThread(c *gin.Context) {
	type JWTToken struct {
		TokenString string `json:"stringToken"`
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
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
		// fmt.Println(currentUser)

		var thread models.Thread

		thread.ID = token.ID
		thread.Name = token.Name
		thread.Description = token.Description
		config.DB.Create(&thread)
		c.JSON(200, thread)
	}

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

	type JWTToken struct {
		TokenString string `json:"stringToken"`
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
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

		var thread, newData models.Thread

		thread.ID = token.ID
		thread.Name = token.Name
		thread.Description = token.Description
		config.DB.Where("id = ?", thread.ID).First(&newData)
		newData.Name = thread.Name
		newData.Description = thread.Description

		config.DB.Save(&newData)
		c.JSON(200, newData)
	}

}

func DeleteThread(c *gin.Context) {

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

		var thread models.Thread
		thread.ID = token.ID
		var count int64
		config.DB.Model(models.Thread{}).Where("id = ?", token.ID).First(&thread).Count(&count)
		if count > 0 {
			config.DB.Delete(&thread)
			c.JSON(200, thread)
			return
		}
		c.JSON(200, "thread not found")
	}

	c.JSON(403, gin.H{"error": "Forbidden"})
	c.Abort()
	c.Redirect(http.StatusTemporaryRedirect, "/")

}
