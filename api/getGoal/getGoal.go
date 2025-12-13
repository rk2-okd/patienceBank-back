package getgoal

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rk2-okd/patienceBank-back/model"
	"gorm.io/gorm"
)

func GetGoalHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		log.Println("GetGoalHandlerにアクセスされました")

		var goal model.Goal
		err := db.Order("created_at desc").First(&goal).Error

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Println("目標データなし → デフォルトGoalを返します")
				goal = model.Goal{
					GoalMoney: 0,
					GoalCount: 0,
				}
				c.JSON(http.StatusOK, goal)
				return
			}

			// 本当のエラーだけ500
			log.Println("GetGoal error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "目標データの取得に失敗しました"})
			return
		}

		// 正常に取得できたとき
		c.JSON(http.StatusOK, goal)
	}
}
