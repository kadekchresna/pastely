package model

import "time"

type Paste struct {
	Shortlink                 string    `json:"shortlink"`
	ExpirationLengthInMinutes int       `json:"expiration_length_in_minutes"`
	PasteURL                  string    `json:"paste_url"`
	CreatedAt                 time.Time `json:"created_at"`
}
