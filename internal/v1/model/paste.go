package model

import "time"

type Paste struct {
	Shortlink                 string    `json:"shortlink"`
	ExpirationLengthInMinutes int       `json:"expiration_length_in_minutes"`
	PasteURL                  string    `json:"paste_url"`
	PasteContent              string    `json:"paste_contents"`
	Status                    string    `json:"status"`
	CreatedAt                 time.Time `json:"created_at"`
	ExpiredAt                 time.Time `json:"expired_at"`
}
