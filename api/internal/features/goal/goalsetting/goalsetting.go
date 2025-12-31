package goalserting

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rk2-okd/patienceBank-back/internal/shared/model"
	"gorm.io/gorm"
)

func GoalSettingsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		db = db.Debug()

		c.Header("Content-Type", "application/json")
		log.Println("GoalSettingsHandlerにアクセスされました")

		var goal model.Goals
		validate := validator.New()

		// JSONを単体としてバインド
		if err := c.ShouldBindJSON(&goal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"バインドエラー": err.Error()})
			return
		}

		// バリデーション
		if err := validate.Struct(goal); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"バリデーションエラー": err.Error()})
			return
		}

		// DB保存
		if err := db.Create(&goal).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"ＤＢ保存エラー": err.Error()})
			return
		}

		// 成功レスポンス
		c.JSON(http.StatusOK, gin.H{
			"message": "目標設定を保存しました",
			"goal_id": goal.GoalID,
		})
	}
}
