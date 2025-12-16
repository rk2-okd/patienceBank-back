package model

import "time"

type Goal struct {
	GoalID    int       `json:"goal_id" gorm:"column:goal_id;primaryKey;autoIncrement"`
	GoalMoney int       `json:"goal_money"`
	GoalCount int       `json:"goal_count"`
	CreatedAt time.Time `json:"created_at"`
}
