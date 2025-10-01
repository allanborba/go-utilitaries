package clock

import "time"

type Clock struct{}

const BRAZILIAN_TIMEZONE = -3

func NewClock() *Clock {
	return &Clock{}
}

func (this *Clock) Now() time.Time {
	return time.Now().UTC()
}

func (this *Clock) NowStringDate() string {
	return this.Now().Format("2006-01-02")
}

func (this *Clock) NowStringDateTime() string {
	return this.Now().Format("2006-01-02 15:04:05")
}

func (this *Clock) NowBrazilianParsed() string {
	return this.Now().Add(BRAZILIAN_TIMEZONE * time.Hour).Format("02/01/2006")
}
