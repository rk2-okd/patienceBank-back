package getgoal

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rk2-okd/patienceBank-back/model"
	"gorm.io/gorm"
)

func GetGoalHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var goal model.Goal
		if err := db.Order("created_at desc").First(&goal).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "目標データの取得に失敗しました"})
			return
		}
		c.JSON(http.StatusOK, goal)
	}
}
