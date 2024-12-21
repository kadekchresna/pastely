package paste

import (
	"time"
)

type Paste struct {
	ID        int64     `json:"id" gorm:"column:id"`
	Shortlink string    `json:"shortlink" gorm:"column:shortlink"`
	PasteURL  string    `json:"paste_url" gorm:"column:paste_url"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	Status    string    `json:"status" gorm:"column:status"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column:expired_at"`
}

func (p Paste) TableName() string {
	return "paste"
}

func (p *Paste) GenerateShortURLBase62() {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	num := time.Now().Unix() + p.ID
	for num > 0 {
		reminder := num % 62
		p.Shortlink = string(alphabet[reminder]) + p.Shortlink
		num = num / 62
	}

}
