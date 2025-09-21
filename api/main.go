package main

import (
	"net/http"

	"github.com/rk2-okd/patienceBank-back/connect"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	connect.DBConnect()
	// Ginのデフォルトのルーターを作成
	r := gin.Default()

	// ルートエンドポイントを設定
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})
	// r.GET("/goalsettings", goalsettings.GoalSettingsHandler(connect.DB))

	// サーバーを起動
	r.Run(":8080") // デフォルトでポート8080で起動
}
