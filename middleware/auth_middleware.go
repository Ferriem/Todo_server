package middleware

import (
	"github.com/Ferriem/Todo_server/models"
	"github.com/Ferriem/Todo_server/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get password from header
		var auth models.Auth
		if err := c.ShouldBindBodyWith(&auth, binding.JSON); err != nil || auth.ID == "" || auth.Password == "" {
			c.JSON(500, gin.H{"error": "Invalid json"})
		}

		if t, err := utils.Login(auth.ID, auth.Password); err != utils.NoError || !t {
			if auth.ID == "" {
				c.JSON(401, gin.H{
					"message": "Please input your username and password",
				})
				c.Abort()
				return
			}
			if err == utils.ErrUserNotFound {
				c.JSON(401, gin.H{
					"message": "User not found, please register first",
				})
			}
			c.JSON(401, gin.H{
				"message": "password error",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
