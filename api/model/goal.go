package model

import "time"

type Goal struct {
	GoalID    int       `json:"goal_id" gorm:"column:goal_id;primaryKey;autoIncrement"`
	Goal      string    `json:"goal" gorm:"column:goal"`
	CreatedAt time.Time `json:"created_at"`
}
