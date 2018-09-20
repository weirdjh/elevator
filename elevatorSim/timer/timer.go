package timer

import (
	"time"
)

// TODO: 2x 4x time

type Timer struct {
	ProgramInitTime time.Time
	CurTime         time.Time
	TimeDiffer      time.Duration
}

func NewTimer() *Timer {

	t1 := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Now().UTC()

	return &Timer{
		ProgramInitTime: t1,
		CurTime:         t2,
		TimeDiffer:      t2.Sub(t1),
	}
}

func (t *Timer) GetTime() time.Time {
	return time.Now().UTC().Add(-t.TimeDiffer)
}

/*
func Log(l *log.Logger, msg string) {
    l.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " [AAA] ")
    l.Print(msg)
}
*/
