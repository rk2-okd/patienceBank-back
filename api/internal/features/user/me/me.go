package me

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// func MeHandler(db *gorm.DB) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		_, err := c.Cookie("token")
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"loggedIn": false,
// 			})
// 			return
// 		}

// 		// セッション検証

//			c.JSON(http.StatusOK, gin.H{
//				"loggedIn": true,
//			})
//		}
//	}
func MeHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// middlewareで必ず入っている前提
		userID := c.MustGet("user_id").(int)

		c.JSON(http.StatusOK, gin.H{
			"loggedIn": true,
			"user_id":  userID, // 必要なら返す
		})
	}
}
