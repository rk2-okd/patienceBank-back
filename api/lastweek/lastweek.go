package lastweek

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rk2-okd/patienceBank-back/model"
	"gorm.io/gorm"
)

func LastWeekHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var records []model.Record
		c.Header("Content-Type", "application/json; charset=utf-8")
		log.Println("LastWeekHandlerにアクセスされました")

		// 今日の日付
		today := time.Now()

		// 今日の曜日（0=日曜, 1=月曜, ..., 6=土曜）
		weekday := int(today.Weekday())

		// 先週の日曜を取得（今日から weekday + 7 日前）
		start := today.AddDate(0, 0, -(weekday + 7))

		// 先週の土曜の翌日（つまり今週の日曜）
		end := start.AddDate(0, 0, 7)

		// DB検索
		if err := db.Where("gaman_day >= ? AND gaman_day < ?", start, end).Find(&records).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "データ取得に失敗しました"})
			return
		}

		// 件数と合計金額を計算
		totalMoney := 0
		for _, record := range records {
			totalMoney += record.GamanMoney // Moneyフィールドがint型であることを想定
		}

		// 結果を返す
		c.JSON(http.StatusOK, gin.H{
			"count": len(records),
			"total": totalMoney,
			"start": start.Format("2006-01-02"),
			"end":   end.AddDate(0, 0, -1).Format("2006-01-02"), // endは今週の日曜なので、土曜にするために1日引く
		})
		log.Printf("取得件数: %d, 合計金額: %d\n", len(records), totalMoney)
	}
}
