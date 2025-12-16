package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	connect "github.com/rk2-okd/patienceBank-back/connect"
	getGoal "github.com/rk2-okd/patienceBank-back/getGoal"
	goalsettings "github.com/rk2-okd/patienceBank-back/goalSettings"
	history "github.com/rk2-okd/patienceBank-back/history"
	"github.com/rk2-okd/patienceBank-back/input"
	lastweek "github.com/rk2-okd/patienceBank-back/lastweek"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := connect.DBConnect()

	// // http://localhost:8080/goalsettings
	// r.POST("/goalsettings", goalsettings.GoalSettingsHandler(connect.DB))
	// goalSettings パッケージの SetupRouter を使う
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: false,
	}))
	r.GET("/getGoal", getGoal.GetGoalHandler(db))
	r.POST("/goalsettings", goalsettings.GoalSettingsHandler(db))
	r.GET("/history", history.HistoryHandler(db))
	r.GET("/lastweek", lastweek.LastWeekHandler(db))
	r.POST("/input", input.InputHandler(db))
	// サーバーを起動
	r.Run(":8080") // デフォルトでポート8080で起動
}
