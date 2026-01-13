package history

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rk2-okd/patienceBank-back/internal/shared/checkuser"
	"github.com/rk2-okd/patienceBank-back/internal/shared/model"
	"github.com/rk2-okd/patienceBank-back/internal/shared/usecase"
	"gorm.io/gorm"
)

func HistoryHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var records []model.Records
		requestRecords := make([]model.RequestRecords, 0, len(records))
		c.Header("Content-Type", "application/json; charset=utf-8")
		log.Println("HistoryHandlerにアクセスされました")
		uid, ok := checkuser.CheckUser(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "ログインしてください"})
			return
		}
		daysParam := c.Query("days")
		if daysParam == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "daysパラメータが必要です"})
			return
		}
		jst, _ := time.LoadLocation("Asia/Tokyo")
		day, err := time.ParseInLocation("2006-01-02", daysParam, jst)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "daysの形式が不正です"})
			return
		}
		start := day
		end := day.AddDate(0, 0, 1)
		if err := db.
			Where("workout_date >= ? AND workout_date < ? AND user_id = ?", start, end, uid).
			Find(&records).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "データ取得に失敗しました"})
			return
		}
		for _, r := range records {
			requestRecords = append(requestRecords, model.RequestRecords{
				WorkoutId:        r.WorkoutId,
				TrainedPart:      usecase.ConvertTrainedPartToString(r.TrainedPart),
				WorkoutDurations: r.WorkoutDurations,
				WorkoutDate:      r.WorkoutDate,
				UserID:           r.UserID,
			})
		}

		log.Printf("取得したレコード: %+v\n", requestRecords) // 取得したレコード数をログに出力

		c.JSON(http.StatusOK, requestRecords)
	}
}
