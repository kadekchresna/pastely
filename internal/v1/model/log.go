package model

import "time"

type Log struct {
	Shortlink string    `json:"shortlink"`
	Total     int       `json:"total"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
}
