package middleware

import (
	"net/http"

	"github.com/PeterHW963/CVWO/backend/models"
	"github.com/gin-gonic/gin"
)

func UserAuthentication(c *gin.Context) {
	// Retrieve the currentUser from the Gin context
	currentUserInterface, exists := c.Get("currentUser")
	if !exists {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	// Type assert the currentUser to your User model
	currentUser, ok := currentUserInterface.(models.User)
	if !ok {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		c.Abort()
		return
	}

	// Now you can access the role field of currentUser
	if currentUser.Role == "User" {
		c.Next()
		return
	}

	c.JSON(403, gin.H{"error": "Forbidden"})
	c.Abort()
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
