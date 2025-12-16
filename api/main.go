package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	connect "github.com/rk2-okd/patienceBank-back/connect"
	goal "github.com/rk2-okd/patienceBank-back/goal"
	record "github.com/rk2-okd/patienceBank-back/record"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := connect.DBConnect()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))
	r.GET("/getGoal", goal.GetGoalHandler(db))
	r.POST("/goalsettings", goal.GoalSettingsHandler(db))
	r.GET("/history", record.HistoryHandler(db))
	r.GET("/lastweek", record.LastWeekHandler(db))
	r.POST("/input", record.InputHandler(db))
	// サーバーを起動
	r.Run(":8080") // デフォルトでポート8080で起動
}
