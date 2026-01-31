package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	getgoal "github.com/rk2-okd/within-back/internal/features/goal/getgoal"
	goalsetting "github.com/rk2-okd/within-back/internal/features/goal/goalsetting"
	input "github.com/rk2-okd/within-back/internal/features/input"
	graph "github.com/rk2-okd/within-back/internal/features/reports/graph"
	history "github.com/rk2-okd/within-back/internal/features/reports/history"
	login "github.com/rk2-okd/within-back/internal/features/user/login"
	logout "github.com/rk2-okd/within-back/internal/features/user/logout"
	me "github.com/rk2-okd/within-back/internal/features/user/me"
	userinfo "github.com/rk2-okd/within-back/internal/features/user/userinfo"
	authmw "github.com/rk2-okd/within-back/internal/shared/auth"
	connect "github.com/rk2-okd/within-back/internal/shared/connect"
)

func main() {
	db := connect.DBConnect()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))
	// ② ★セッション（cookie）を使う設定：これが「ログイン状態を保持する土台」
	store := cookie.NewStore([]byte("super-secret-key")) // 本番は環境変数にする
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 1, // 1日
		HttpOnly: true,
		Secure:   false, // https のときだけ true
		// SameSite: http.SameSiteLaxMode, // 必要なら設定（多くはデフォルトでOK）
	})
	r.Use(sessions.Sessions("within_session", store))
	r.POST("/login", login.LoginHandler(db))
	auth := r.Group("/")
	auth.Use(authmw.AuthRequired())
	{
		auth.GET("/me", me.MeHandler(db))
		auth.GET("/history", history.HistoryHandler(db))
		auth.POST("/input", input.InputHandler(db))
		auth.GET("/graph", graph.GraphHandler(db))
		auth.GET("/getgoal", getgoal.GetGoalHandler(db))
		auth.GET("/getUser", userinfo.GetUserHandler(db))
		auth.POST("/logout", logout.LogoutHandler())
		auth.PATCH("/goalsetting", goalsetting.GoalSettingsHandler(db))
	}
	for _, rt := range r.Routes() {
		println("ROUTE:", rt.Method, rt.Path)
	}
	r.Run(":8080")
}
