package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/PeterHW963/CVWO/backend/config"
	"github.com/PeterHW963/CVWO/backend/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

// func UserController(c *gin.Context) {
// 	c.String(200, "Hello World!")
// }

// func AllUser(c *gin.Context) {
// 	users := []models.User{}
// 	config.DB.Find(&users)
// 	fmt.Println(users)
// 	c.JSON(200, users)

// }

// search for user (on search bar)

func GetUser(c *gin.Context) {
	var data struct {
		Search string `json:"search"`
	}

	c.ShouldBindJSON(&data)
	users := []models.User{}
	config.DB.Where("name LIKE ?", "%"+data.Search+"%").Find(&users)
	fmt.Println(users)
	c.JSON(200, users)

}

// create user

func CreateUser(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)

	var count int64 = 0
	config.DB.Model(models.User{}).Where("email=?", user.Email).Count(&count)
	if count != 0 {
		c.String(200, "email already used")
		return
	}

	config.DB.Model(models.User{}).Where("username=?", user.UserName).Count(&count)
	if count != 0 {
		c.String(200, "username already used")
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)

	if err != nil {
		c.String(200, "Hash")
		return
	}

	user.Password = string(hashedPassword)
	config.DB.Create(&user)
	c.JSON(200, "User created")

}

func Login(c *gin.Context) {
	var loginUser, user models.User
	c.ShouldBindJSON(&loginUser)

	resultUser := config.DB.First(&user, "user_name=?", loginUser.UserName)
	if resultUser.RowsAffected == 0 {
		c.String(200, "user not found")
		return
	}

	resultPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if resultPassword != nil {
		c.String(200, "wrong password")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.ID,
		"expires": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("KEY")))
	if err != nil {
		c.String(200, "Generating JWT Failed")
		return
	}

	c.String(200, tokenString)
}

func Authenticate(c *gin.Context) {
	currentUser, _ := c.Get("currentUser")
	c.JSON(200, currentUser)
}

// delete user
func DeleteUser(c *gin.Context) {

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

		var user models.User
		user.ID = token.ID
		var count int64
		config.DB.Model(models.User{}).Where("id = ?", token.ID).First(&user).Count(&count)
		if count > 0 {
			config.DB.Delete(&user)
			c.JSON(200, user)
			return
		}
		c.JSON(200, "user not found")
	}

	c.JSON(403, gin.H{"error": "Forbidden"})
	c.Abort()
	c.Redirect(http.StatusTemporaryRedirect, "/")

}

func UpdateUser(c *gin.Context) {

	type JWTToken struct {
		TokenString string `json:"stringToken"`
		ID          uint   `json:"id"`
		Username    string `json:"username"`
		Email       string `json:"email"`
		Password    string `json:"password"`
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

		var user, newData models.User
		user.ID = token.ID
		user.UserName = token.Username
		user.Email = token.Email
		user.Password = token.Password

		config.DB.Where("id = ?", user.ID).First(&newData)

		if currentUser.ID != newData.ID {
			c.JSON(403, gin.H{"error": "Forbidden: You can only update your own posts"})
			return
		}

		if newData.UserName != user.UserName && user.UserName != "" {
			newData.UserName = user.UserName
		} else if newData.Email != user.Email && user.Email != "" {
			newData.Email = user.Email
		} else if newData.Password != user.Password && user.Password != "" {
			newData.Password = user.Password
		}
		config.DB.Save(&newData)
		c.JSON(200, newData)
	}

}
