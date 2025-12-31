package login

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rk2-okd/patienceBank-back/internal/shared/model"
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
		// ③ トークン発行（ここでは例としてダミー。実運用はJWTを作る）
		//    JWTにする場合は user.ID を入れて署名して返す
		token := "dummy-token-for-example"
		// ④ httpOnly Cookie にセット
		// 開発(localhost)は Secure=false（httpsじゃないから）
		// SameSite は c.SetSameSite で指定できる（後述）
		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie(
			"token",
			token,
			int((24 * time.Hour).Seconds()), // 1日
			"/",
			"",
			false, // httpsなら true
			true,  // httpOnly
		)
		c.JSON(http.StatusOK, gin.H{
			"message":  "login success",
			"username": user.Username,
		})
	}
}
