package log

import "time"

type GetLogParams struct {
	Shortlink string    `query:"shortlink"`
	DateStart time.Time `query:"date_start"`
	DateEnd   time.Time `query:"date_end"`
}

func NewGetLogParams(Shortlink string, DateStart time.Time, DateEnd time.Time) GetLogParams {
	return GetLogParams{
		Shortlink: Shortlink,
		DateStart: DateStart,
		DateEnd:   DateEnd,
	}
}

type CreateLog struct {
	Shortlink string `query:"shortlink"`
}

func NewCreateLog(Shortlink string) CreateLog {
	return CreateLog{
		Shortlink: Shortlink,
	}
}
