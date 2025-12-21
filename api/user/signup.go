package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/rk2-okd/patienceBank-back/model"
)

type SignupRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func SignupHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req SignupRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid request",
			})
			return
		}

		// パスワードをハッシュ化
		hash, err := bcrypt.GenerateFromPassword(
			[]byte(req.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "hash error",
			})
			return
		}

		user := model.User{
			Username:     req.Username,
			Email:        req.Email,
			PasswordHash: string(hash),
		}

		// users に INSERT
		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "user already exists",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "signup success",
		})
	}
}
