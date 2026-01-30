package model

type Users struct {
	ID           *int   `json:"id" gorm:"column:id"`
	Username     string `json:"username" gorm:"column:username"`
	Email        string `json:"email" gorm:"column:email"`
	PasswordHash string `json:"password_hash" gorm:"column:password_hash"`
	Goal         string `json:"goal" gorm:"column:goal"`
	AuthProvider string `json:"auth_provider" gorm:"column:auth_provider"`
}
