package history

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rk2-okd/patienceBank-back/internal/shared/model"
	"gorm.io/gorm"
)

func HistoryHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var records []model.Records
		c.Header("Content-Type", "application/json; charset=utf-8")

		daysParam := c.Query("days")
		if daysParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "daysパラメータが必要です"})
			return
		}

		// DATE列ならこれが一番確実
		if err := db.Where("workout_date = ?", daysParam).Find(&records).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "データ取得に失敗しました"})
			return
		}

		c.JSON(http.StatusOK, records)
	}
}
