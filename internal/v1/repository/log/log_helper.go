package log

import "time"

type GetLogParams struct {
	Shortlink string
	DateStart time.Time
	DateEnd   time.Time
}

func NewGetLogParams(Shortlink string, DateStart time.Time, DateEnd time.Time) GetLogParams {
	return GetLogParams{
		Shortlink: Shortlink,
		DateStart: DateStart,
		DateEnd:   DateEnd,
	}
}
