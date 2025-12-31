package input

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rk2-okd/patienceBank-back/internal/shared/model"
	"gorm.io/gorm"
)

func InputHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		log.Println("InputHandlerにアクセスされました")
		var record model.Records
		validate := validator.New()
		if err := c.ShouldBindJSON(&record); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"バインドエラー": err.Error()})
			return
		}
		log.Printf("bind result: %+v\n", record)
		if err := validate.Struct(record); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"バリデーションエラー": err.Error()})
			return
		}
		record.UserID = 1
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
