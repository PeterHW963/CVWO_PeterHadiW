package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/PeterHW963/CVWO/backend/config"
	"github.com/PeterHW963/CVWO/backend/models"
	"github.com/gin-gonic/gin"
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
	var data struct {
		ID int `json:"id"`
	}

	c.ShouldBindJSON(&data)
	var user models.User
	var count int64
	config.DB.Model(models.User{}).Where("id =?", data.ID).First(&user).Count(&count)
	if count > 0 {
		config.DB.Delete(&user)
		c.JSON(200, user)
		return
	}
	c.JSON(200, "user not found")
}

func UpdateUser(c *gin.Context) {
	var user, newData models.User
	if err := c.ShouldBindJSON(&user); err != nil {
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

	if currentUser.ID != user.ID {
		c.JSON(403, gin.H{"error": "Forbidden"})
		c.Abort()
		return
	}
	config.DB.Where("id =?", user.ID).First(&newData)
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
