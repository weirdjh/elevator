package timer

import (
	"time"
)

// TODO: 2x 4x time

type Timer struct {
	curTime time.Time
}

func (t *Timer) NewTimer() *Timer {
	return &Timer{
		curTime: time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
}
