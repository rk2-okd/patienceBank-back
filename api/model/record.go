package model

type Record struct {
	GamanID      int    `json:"gaman_id" gorm:"column:gaman_id;primaryKey;autoIncrement"`
	GamanThing   string `json:"gaman_thing" gorm:"column:gaman_thing"`
	GamanMinutes int    `json:"gaman_minutes" gorm:"column:gaman_minutes"`
	GamanDay     string `json:"gaman_day" gorm:"column:gaman_day"`
	UserID       int    `json:"user_id" gorm:"column:user_id"`
}

func (Record) TableName() string { return "records" }
