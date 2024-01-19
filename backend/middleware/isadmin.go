package middleware

import (
	"net/http"

	"github.com/PeterHW963/CVWO/backend/models"
	"github.com/gin-gonic/gin"
)

func AdminAuthentication(c *gin.Context) {
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

	if currentUser.Role == "Admin" {
		c.Next()
		return
	}

	c.JSON(403, gin.H{"error": "Forbidden"})
	c.Abort()
	c.Redirect(http.StatusTemporaryRedirect, "/")
}
