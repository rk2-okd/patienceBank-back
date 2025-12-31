package model

type Users struct {
	ID           uint   `json:"id" gorm:"column:id"`
	Username     string `json:"username" gorm:"column:username"`
	Email        string `json:"email" gorm:"column:email"`
	PasswordHash string `json:"password_hash" gorm:"column:password_hash"`
	Comment      string `json:"comment" gorm:"column:comment"`
	AuthProvider string `json:"auth_provider" gorm:"column:auth_provider"`
}
