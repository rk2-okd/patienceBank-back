package model

import "time"

type Goals struct {
	GoalID    int       `json:"goal_id" gorm:"column:goal_id"`
	UserID    int       `json:"user_id" gorm:"column:user_id"`
	Goal      string    `json:"goal" gorm:"column:goal"`
	CreatedAt time.Time `json:"created_at"`
}
