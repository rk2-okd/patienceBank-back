package input

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rk2-okd/patienceBank-back/internal/shared/checkuser"
	"github.com/rk2-okd/patienceBank-back/internal/shared/model"
	"github.com/rk2-okd/patienceBank-back/internal/shared/usecase"
	"gorm.io/gorm"
)

func InputHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		log.Println("InputHandlerにアクセスされました")
		var record model.Records
		var req model.RequestRecords
		uid, ok := checkuser.CheckUser(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "ログインしてください"})
			return
		}
		validate := validator.New()
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"バインドエラー": err.Error()})
			return
		}
		jst, _ := time.LoadLocation("Asia/Tokyo")
		record = model.Records{
			TrainedPart:      usecase.ConvertTrainedPartToTinyInt(req.TrainedPart),
			WorkoutDurations: req.WorkoutDurations,
			WorkoutDate:      time.Now().In(jst),
			UserID:           uid,
		}
		log.Printf("bind result: %+v\n", record)
		if err := validate.Struct(record); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"バリデーションエラー": err.Error()})
			return
		}
		if err := db.Create(&record).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"ＤＢ保存エラー": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":   "記録を保存しました",
			"record_id": record.WorkoutId,
		})
	}
}
