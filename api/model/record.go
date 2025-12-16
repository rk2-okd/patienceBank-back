package model

type Record struct {
	GamanID    int    `json:"gaman_id" gorm:"column:gaman_id;primaryKey;autoIncrement"`
	GamanThing string `json:"gaman_thing" gorm:"column:gaman_thing"`
	GamanMoney int    `json:"gaman_money" gorm:"column:gaman_money"`
	GamanDay   string `json:"gaman_day" gorm:"column:gaman_day"`
}

func (Record) TableName() string { return "records" }
