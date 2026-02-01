package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

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

	_ "github.com/go-sql-driver/mysql"
	connect "github.com/rk2-okd/within-back/internal/shared/connect"
)

func main() {
	// ローカル用: .env を読む（本番は .env が無いのが普通なので Fatal にしない）
	if err := godotenv.Load(); err != nil {
		log.Println(".env not loaded (this is OK on AWS):", err)
	}
	db := connect.DBConnect()
	r := gin.Default()

	// -----------------------------
	// 1) CORS: 環境変数から許可Originを読む
	// -----------------------------
	origins := getOriginsFromEnv("FRONTEND_ORIGINS",
		[]string{"http://localhost:3000"},
	)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	// -----------------------------
	// 2) セッション鍵: 環境変数から読む（必須）
	// -----------------------------
	secret := os.Getenv("SESSION_SECRET")
	if len(secret) < 32 {
		// 32未満だと弱い。未設定もここで止める（本番事故防止）
		log.Fatal("SESSION_SECRET is missing or too short (need 32+ chars)")
	}

	store := cookie.NewStore([]byte(secret))

	// -----------------------------
	// 3) CookieのHTTPS対応
	// -----------------------------
	// 本番でhttps運用するなら true にする（環境変数で切り替え）
	cookieSecure := os.Getenv("COOKIE_SECURE") == "true"

	// フロントとAPIが別ドメインで、cookieでログイン維持したいなら None が必要なことが多い
	// None を使う場合は Secure=true が必須（ブラウザ仕様）
	sameSite := http.SameSiteLaxMode
	if os.Getenv("COOKIE_SAMESITE_NONE") == "true" {
		sameSite = http.SameSiteNoneMode
		if !cookieSecure {
			log.Fatal("COOKIE_SAMESITE_NONE=true requires COOKIE_SECURE=true")
		}
	}

	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 1, // 1日
		HttpOnly: true,
		Secure:   cookieSecure,
		SameSite: sameSite,
	})

	r.Use(sessions.Sessions("within_session", store))

	// routes
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

	// -----------------------------
	// 4) PORT: 環境変数があればそれを使う（EB対策）
	// -----------------------------
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Listening on port", port)
	r.Run(":" + port)

}

func getOriginsFromEnv(key string, fallback []string) []string {
	v := strings.TrimSpace(os.Getenv(key))
	if v == "" {
		return fallback
	}
	parts := strings.Split(v, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		s := strings.TrimSpace(p)
		if s != "" {
			out = append(out, s)
		}
	}
	if len(out) == 0 {
		return fallback
	}
	return out
}
