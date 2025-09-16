package main

import (
	"net/http"

	"github.com/patience-back/api/connect"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connect.DBconnect()
	// Ginのデフォルトのルーターを作成
	r := gin.Default()

	// ルートエンドポイントを設定
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	r.GET("/goalsettings", func(c *gin.Context) {

	})

	// サーバーを起動
	r.Run(":8080") // デフォルトでポート8080で起動
}
