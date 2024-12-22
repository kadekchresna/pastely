package model

import "time"

type Log struct {
	Time      time.Time `json:"time"`
	Shortlink string    `json:"shortlink"`
	Total     int       `json:"total"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
}
