package model

import "time"

type Records struct {
	WorkoutId        int       `json:"workout_id" gorm:"column:workout_id"`
	TrainedPart      int       `json:"trained_part" gorm:"column:trained_part"`
	WorkoutDurations int       `json:"workout_duration" gorm:"column:workout_duration"`
	WorkoutDate      time.Time `json:"workout_date" gorm:"column:workout_date"`
	UserID           int       `json:"user_id" gorm:"column:user_id"`
}
