package goalsetting

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rk2-okd/within-back/internal/shared/checkuser"
	"github.com/rk2-okd/within-back/internal/shared/model"
	"gorm.io/gorm"
)

func GoalSettingsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		db = db.Debug()

		c.Header("Content-Type", "application/json")
		log.Println("GoalSettingsHandlerにアクセスされました")

		uid, ok := checkuser.CheckUser(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "ログインしてください"})
			return
		}
		// ===== ここから追加 =====
		raw, _ := c.GetRawData()
		log.Printf("raw body: %q", string(raw))
		c.Request.Body = io.NopCloser(strings.NewReader(string(raw)))
		// ===== ここまで追加 =====
		var req model.GoalSettingReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"バインドエラー": err.Error()})
			return
		}
		goal := strings.TrimSpace(req.Goal)
		if goal == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "goal は必須です"})
			return
		}
		if err := db.Model(&model.Users{}).
			Where("id = ?", uid).
			Update("goal", goal).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"ＤＢ更新エラー": err.Error()})
			return
		}

		// 成功レスポンス
		c.JSON(http.StatusOK, gin.H{
			"message": "目標設定を更新しました",
		})
	}
}
