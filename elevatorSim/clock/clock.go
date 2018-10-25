package clock

import (
	"time"
)

type Clock struct {
	VirtualTime  time.Time
	RealTime     time.Time
	RealStopTime time.Time
	Factor       int
}

func NewClock() *Clock {

	t1 := time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC)
	t2 := time.Now()

	return &Clock{
		VirtualTime: t1,
		RealTime:    t2,
		Factor:      1,
	}
}

func (c *Clock) GetClock() time.Time {
	return c.VirtualTime.Add(time.Now().Sub(c.RealTime) * time.Duration(c.Factor))
}

func (c *Clock) StartClock() {
	c.Factor = 1
}

func (c *Clock) StopClock() {
	c.RealStopTime = time.Now().UTC()
}

func (c *Clock) FasterClock() {
	c.Factor *= 4
}

/*
func Log(l *log.Logger, msg string) {
    l.SetPrefix(time.Now().Format("2006-01-02 15:04:05") + " [AAA] ")
    l.Print(msg)
}
*/
