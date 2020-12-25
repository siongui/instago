package igdl

import (
	"time"
)

type TimeLimiter struct {
	LastTime time.Time
	Interval int64 // unit: second
}

func NewTimeLimiter(interval int64) *TimeLimiter {
	tl := TimeLimiter{
		LastTime: time.Unix(time.Now().Unix()-interval, 0),
		Interval: interval,
	}
	return &tl
}

func (tl *TimeLimiter) WaitAtLeastIntervalAfterLastTime() {
	d := time.Now().Sub(tl.LastTime)
	interval := time.Duration(tl.Interval) * time.Second
	if d < interval {
		time.Sleep(interval - d)
	}
}

func (tl *TimeLimiter) SetLastTimeToNow() {
	tl.LastTime = time.Now()
}
