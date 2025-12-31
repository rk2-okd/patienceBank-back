package graph

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rk2-okd/patienceBank-back/internal/shared/model"
	"gorm.io/gorm"
)

func GraphHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var records []model.Records
		c.Header("Content-Type", "application/json; charset=utf-8")
		log.Println("GraphHandlerにアクセスされました")

		if err := db.Find(&records).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "データ取得に失敗しました"})
			return
		}
		c.JSON(http.StatusOK, records)
		log.Println("取得したレコード:", records) // 取得したレコード数をログに出力
	}
}
