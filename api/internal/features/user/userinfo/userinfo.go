package userinfo

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rk2-okd/within-back/internal/shared/model"
	"gorm.io/gorm"
)

func GetUserHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("GetUserHandlerにアクセスされました")

		userID := c.Param("id")

		var user model.Users
		if err := db.First(&user, userID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーの取得に失敗しました"})
			return
		}

		// レスポンス用に安全な形にする
		c.JSON(http.StatusOK, gin.H{
			"id":            user.ID,
			"username":      user.Username,
			"email":         user.Email,
			"goal":          user.Goal,
			"auth_provider": user.AuthProvider,
		})
	}
}
