package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	connect "github.com/rk2-okd/patienceBank-back/connect"
	getGoal "github.com/rk2-okd/patienceBank-back/getGoal"
	goalsettings "github.com/rk2-okd/patienceBank-back/goalSettings"
	"github.com/rk2-okd/patienceBank-back/history"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := connect.DBConnect()

	// // http://localhost:8080/goalsettings
	// r.POST("/goalsettings", goalsettings.GoalSettingsHandler(connect.DB))
	// goalSettings パッケージの SetupRouter を使う
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},            // フロントエンドのオリジンを指定
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},           // 許可するメソッド
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"}, // 許可するヘッダー
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,           // クッキーや認証情報を許可（必要に応じて）
		MaxAge:           12 * time.Hour, // プリフライトリクエストのキャッシュ時間
	}))
	r.GET("/getGoal", getGoal.GetGoalHandler(db))
	r.POST("/goalsettings", goalsettings.GoalSettingsHandler(db))
	r.GET("/history", history.HistoryHandler(db))

	// サーバーを起動
	r.Run(":8080") // デフォルトでポート8080で起動
}
