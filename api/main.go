package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	getgoal "github.com/rk2-okd/patienceBank-back/internal/features/goal/getgoal"
	goalsetting "github.com/rk2-okd/patienceBank-back/internal/features/goal/goalsetting"
	input "github.com/rk2-okd/patienceBank-back/internal/features/input"
	graph "github.com/rk2-okd/patienceBank-back/internal/features/reports/graph"
	history "github.com/rk2-okd/patienceBank-back/internal/features/reports/history"
	login "github.com/rk2-okd/patienceBank-back/internal/features/user/login"
	me "github.com/rk2-okd/patienceBank-back/internal/features/user/me"
	userinfo "github.com/rk2-okd/patienceBank-back/internal/features/user/userinfo"

	// lastweek "github.com/rk2-okd/patienceBank-back/internal/features/reports/lastweek"
	connect "github.com/rk2-okd/patienceBank-back/internal/shared/connect"

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
	r.GET("/getGoal", getgoal.GetGoalHandler(db))
	r.POST("/goalsettings", goalsetting.GoalSettingsHandler(db))
	r.GET("/history", history.HistoryHandler(db))
	r.GET("/graph", graph.GraphHandler(db))
	// r.GET("/lastweek", reports.LastWeekHandler(db))
	r.POST("/input", input.InputHandler(db))
	r.POST("/login", login.LoginHandler(db))
	r.GET("/me", me.MeHandler(db))
	r.GET("/getUser", userinfo.GetUserHandler(db))

	// サーバーを起動
	r.Run(":8080") // デフォルトでポート8080で起動
}
