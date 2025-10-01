package clock

import "time"

type ClockFake struct {
	NowReturn time.Time
}

func NewClockFake() *ClockFake {
	return &ClockFake{NowReturn: time.Date(2025, 6, 12, 14, 24, 56, 0, time.UTC)}
}

func (this *ClockFake) Now() time.Time {
	return this.NowReturn
}

func (this *ClockFake) NowStringDate() string {
	return this.Now().Format("2006-01-02")
}

func (this *ClockFake) NowStringDateTime() string {
	return this.Now().Format("2006-01-02 15:04:05")
}

func (this *ClockFake) NowBrazilianParsed() string {
	return this.Now().Add(BRAZILIAN_TIMEZONE * time.Hour).Format("02/01/2006")
}
