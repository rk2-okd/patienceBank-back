package history

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rk2-okd/patienceBank-back/model"
	"gorm.io/gorm"
)

func HistoryHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var records []model.Record
		c.Header("Content-Type", "application/json; charset=utf-8")
		log.Println("HistoryHandlerにアクセスされました")
		params := c.Request.URL.Query()
		daysParam := params.Get("days")
		if daysParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "daysパラメータが必要です"})
			return
		}
		parsedDate, err := time.Parse("2006-01-02", daysParam)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "daysパラメータの形式が不正です"})
			return
		}
		start := parsedDate
		end := parsedDate.AddDate(0, 0, 1)

		if err := db.Where("gaman_day >= ? AND gaman_day < ?", start, end).Find(&records).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "データ取得に失敗しました"})
			return
		}
		c.JSON(http.StatusOK, records)
		log.Println("取得したレコード:", records) // 取得したレコード数をログに出力
	}
}
