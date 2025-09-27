package model

import "time"

type Goal struct {
	GoalID    int       `gorm:"primaryKey;autoIncrement" json:"goal_id"`
	GoalMoney int       `json:"goal_money" validate:"required"`
	GoalCount int       `json:"goal_count" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}
