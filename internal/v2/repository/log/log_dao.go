package log

import "time"

type Log struct {
	Time      time.Time `json:"time" gorm:"column:time"`
	Shortlink string    `json:"shortlink" gorm:"column:shortlink"`
}

func (l Log) TableName() string {
	return "paste_log"
}

type LogAggregate struct {
	Shortlink string `json:"shortlink" gorm:"column:shortlink"`
	Total     int    `json:"total" gorm:"column:total"`
}
