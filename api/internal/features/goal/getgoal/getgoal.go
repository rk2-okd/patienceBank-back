package getgoal

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rk2-okd/patienceBank-back/internal/shared/checkuser"
	"github.com/rk2-okd/patienceBank-back/internal/shared/model"
	"gorm.io/gorm"
)

func GetGoalHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		log.Println("GetGoalHandlerにアクセスされました")
		uid, ok := checkuser.CheckUser(c)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "ログインしてください"})
			return
		}
		var user model.Users
		err := db.Where("id = ?", uid).First(&user).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Println("ユーザー情報なし → デフォルトユーザーを返します")
				return
			}
			// 本当のエラーだけ500
			log.Println("GetGoal error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "目標データの取得に失敗しました"})
			return
		}
		// 正常に取得できたとき
		c.JSON(http.StatusOK, user)
	}
}
