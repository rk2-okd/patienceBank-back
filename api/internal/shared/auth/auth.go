package auth

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		v := session.Get("user_id")

		// 未ログイン（セッションに入ってない）
		if v == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "ログインしてください",
			})
			return
		}

		// LoginHandler で int を入れた前提なので、ここでは int を期待する
		id, ok := v.(int)
		if !ok {
			// 型が想定外 = セッションが壊れてる/古い/別の値が入ってる
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "ログインしてください",
			})
			return
		}

		// 後続のhandlerで使えるようにする
		c.Set("user_id", id)
		c.Next()
	}
}
