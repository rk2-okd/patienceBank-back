package model

type Record struct {
	GamanID    int    `json:"gaman_id" varlidate:"required"`
	GamanThing string `json:"gaman_thing" validate:"required"`
	GamanMoney int    `json:"gaman_money" validate:"required"`
	GamanDay   string `json:"gaman_day" validate:"required"`
}
