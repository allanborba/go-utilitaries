package clock

import "time"

var timerInstance ClockInterface

type ClockInterface interface {
	Now() time.Time
	NowStringDate() string
	NowStringDateTime() string
	NowBrazilianParsed() string
}

func GetInstance() ClockInterface {
	return timerInstance
}

func SetTimerInstance(instance ClockInterface) {
	timerInstance = instance
}
