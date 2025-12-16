package user

import (
	"github.com/gin-gonic/gin"
)

func MeHandler(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"loggedIn": false})
		return
	}

	// 本来は token 検証する
	_ = token

	c.JSON(200, gin.H{"loggedIn": true})
}
