package controller

// get post
// create post
// delete post
// udpate post

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

func CreatePost(c *gin.Context) {
	type JWTToken struct {
		TokenString string `json:"stringToken"`
		ID          uint   `json:"id"`
		UserId      uint   `json:"userid"`
		Title       string `json:"title"`
		Content     string `json:"content"`
		TagId       *uint  `json:"tagid"`
		ThreadId    *uint  `json:"threadid"`
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

		// Check if TagId exists
		if token.TagId != nil {
			var tagCount int64
			var tag models.Tag
			config.DB.First(&tag, "id = ?", *token.TagId).Count(&tagCount)
			if tagCount == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Tag not found"})
				return
			}
		}

		// Check if ThreadId exists
		if token.ThreadId != nil {
			var threadCount int64
			var thread models.Thread
			config.DB.First(&thread, "id = ?", *token.ThreadId).Count(&threadCount)
			if threadCount == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Thread not found"})
				return
			}
		}
		var post models.Post

		post.ID = token.ID
		post.UserId = token.UserId
		post.Title = token.Title
		post.Content = token.Content
		post.TagId = token.TagId
		post.ThreadId = token.ThreadId
		config.DB.Create(&post)
		c.JSON(200, post)
	}
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

	type JWTToken struct {
		TokenString string `json:"stringToken"`
		ID          uint   `json:"id"`
		Title       string `json:"title"`
		Content     string `json:"content"`
		TagId       *uint  `json:"tagid"`
		ThreadId    *uint  `json:"threadid"`
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

		var post, newData models.Post

		post.ID = token.ID
		post.Title = token.Title
		post.Content = token.Content
		post.TagId = token.TagId
		post.ThreadId = token.ThreadId

		config.DB.Where("id = ?", post.ID).First(&newData)

		if currentUser.ID != newData.UserId {
			c.JSON(403, gin.H{"error": "Forbidden: You can only update your own posts"})
			return
		}

		// Check if TagId exists
		if token.TagId != nil {
			var tagCount int64
			var tag models.Tag
			config.DB.First(&tag, "id = ?", *token.TagId).Count(&tagCount)
			if tagCount == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Tag not found"})
				return
			}
		}

		// validation of ThreadId existence
		if token.ThreadId != nil {
			var threadCount int64
			var thread models.Thread
			config.DB.First(&thread, "id = ?", *token.ThreadId).Count(&threadCount)
			if threadCount == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Thread not found"})
				return
			}
		}

		newData.Title = post.Title
		newData.Content = post.Content
		config.DB.Save(&newData)
		c.JSON(200, newData)
	}

}

func DeletePost(c *gin.Context) {

	type JWTToken struct {
		TokenString string `json:"stringToken"`
		ID          uint   `json:"id"`
	}

	var token JWTToken
	if err := c.ShouldBindBodyWith(&token, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validation that token is input
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

	// validation that the user is authenticated with a valid token
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

	// validation of admin role
	if currentUser.Role == "Admin" {

		var post models.Post
		post.ID = token.ID
		var count int64
		config.DB.Model(models.Post{}).Where("id = ?", token.ID).First(&post).Count(&count)
		if count > 0 {
			config.DB.Delete(&post)
			c.JSON(200, post)
			return
		}
		c.JSON(200, "post not found")
	}

	c.JSON(403, gin.H{"error": "Forbidden"})
	c.Abort()
	c.Redirect(http.StatusTemporaryRedirect, "/")

}
