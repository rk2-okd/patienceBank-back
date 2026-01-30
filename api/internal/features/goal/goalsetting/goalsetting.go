package goalsetting

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rk2-okd/patienceBank-back/internal/shared/checkuser"
	"github.com/rk2-okd/patienceBank-back/internal/shared/model"
	"gorm.io/gorm"
)

func GoalSettingsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		db = db.Debug()

		c.Header("Content-Type", "application/json")
		log.Println("GoalSettingsHandlerにアクセスされました")

		var user model.Users
		validate := validator.New()
		uid, ok := checkuser.CheckUser(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "ログインしてください"})
			return
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"バインドエラー": err.Error()})
			return
		}
		user = model.Users{
			ID:           &uid,
			Username:     user.Username,
			Email:        user.Email,
			PasswordHash: user.PasswordHash,
			Goal:         strings.TrimSpace(user.Goal),
			AuthProvider: user.AuthProvider,
		}
		// バリデーション
		if err := validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"バリデーションエラー": err.Error()})
			return
		}
		// DB保存
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"ＤＢ保存エラー": err.Error()})
			return
		}

		// 成功レスポンス
		c.JSON(http.StatusOK, gin.H{
			"message": "目標設定を保存しました",
		})
	}
}
