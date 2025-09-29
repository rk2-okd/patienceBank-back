package getgoal

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rk2-okd/patienceBank-back/model"
	"gorm.io/gorm"
)

func GetGoalHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		log.Println("GetGoalHandlerにアクセスされました")
		// DBから最新の目標を取得
		var goal model.Goal
		if err := db.Order("created_at desc").First(&goal).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "目標データの取得に失敗しました"})
			return
		}
		c.JSON(http.StatusOK, goal)
	}
}
