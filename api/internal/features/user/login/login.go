package login

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rk2-okd/within-back/internal/shared/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "入力が不正です"})
			return
		}
		// ① メールアドレスでユーザー取得
		var user model.Users
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			// ユーザーがいない / DBエラーでも、攻撃対策で同じメッセージにするのが無難
			c.JSON(http.StatusUnauthorized, gin.H{"message": "メールアドレスまたはパスワードが違います"})
			return
		}
		// ② パスワード検証（平文req.Password vs DBのハッシュ）
		if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "メールアドレスまたはパスワードが違います"})
			return
		}
		if user.ID == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "ユーザーIDが不正です"})
			return
		}

		session := sessions.Default(c)
		session.Set("user_id", *user.ID) // ← ここがポイント（ポインタを外して int を入れる）
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "セッション保存に失敗しました"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "login success",
			"username": user.Username,
		})
	}
}
