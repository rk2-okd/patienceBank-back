package user

import (
	"github.com/gin-gonic/gin"
)

// POST /login
func LoginHandler(c *gin.Context) {
	// ① ID/Pass 検証（省略）

	token := "jwt-string" // ← 本来はJWTを生成

	c.SetCookie(
		"token",
		token,
		60*60*24, // 1日
		"/",
		"localhost", // 本番ではドメイン
		false,       // httpsなら true
		true,        // httpOnly
	)

	c.JSON(200, gin.H{
		"message": "login success",
	})
}
