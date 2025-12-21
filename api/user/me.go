package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MeHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"loggedIn": false,
			})
			return
		}

		// 本来はここで token を検証する

		c.JSON(http.StatusOK, gin.H{
			"loggedIn": true,
		})
	}
}
