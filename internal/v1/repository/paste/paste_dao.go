package paste

import "time"

type Paste struct {
	Shortlink                 string    `json:"shortlink" gorm:"column:shortlink"`
	ExpirationLengthInMinutes int       `json:"expiration_length_in_minutes" gorm:"column:expiration_length_in_minutes"`
	PasteURL                  string    `json:"paste_url" gorm:"column:paste_url"`
	CreatedAt                 time.Time `json:"created_at" gorm:"column:created_at"`
}

func (p Paste) TableName() string {
	return "paste"
}
