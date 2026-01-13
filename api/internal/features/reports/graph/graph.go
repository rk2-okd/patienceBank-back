package graph

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rk2-okd/patienceBank-back/internal/shared/checkuser"
	"github.com/rk2-okd/patienceBank-back/internal/shared/model"
	"github.com/rk2-okd/patienceBank-back/internal/shared/usecase"
	"gorm.io/gorm"
)

func GraphHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var records []model.Records
		var req []model.RequestRecords
		c.Header("Content-Type", "application/json; charset=utf-8")
		log.Println("GraphHandlerにアクセスされました")
		uid, ok := checkuser.CheckUser(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "ログインしてください"})
			return
		}
		if err := db.
			Where("user_id = ?", uid).
			Find(&records).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "データ取得に失敗しました"})
			return
		}
		for _, r := range records {
			req = append(req, model.RequestRecords{
				WorkoutId:        r.WorkoutId,
				TrainedPart:      usecase.ConvertTrainedPartToString(r.TrainedPart),
				WorkoutDurations: r.WorkoutDurations,
				WorkoutDate:      r.WorkoutDate,
				UserID:           r.UserID,
			})
		}
		c.JSON(http.StatusOK, req)
		log.Println("取得したレコード:", req) // 取得したレコード数をログに出力
	}
}
