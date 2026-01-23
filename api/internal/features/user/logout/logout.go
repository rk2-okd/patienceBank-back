package logout

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		// セッション内データを削除
		session.Clear()

		// Cookie を即時失効（今のセッションを切る）
		session.Options(sessions.Options{
			Path:   "/",
			MaxAge: -1,
		})

		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "ログアウトに失敗しました",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "logout success",
		})
	}
}
