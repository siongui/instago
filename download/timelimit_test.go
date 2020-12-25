package igdl

import (
	"testing"
	"time"
)

func TestTimeLimiter(t *testing.T) {
	tl := NewTimeLimiter(2)
	t.Log(time.Now())
	tl.WaitAtLeastIntervalAfterLastTime()
	t.Log(time.Now())
	tl.SetLastTimeToNow()
	tl.WaitAtLeastIntervalAfterLastTime()
	t.Log(time.Now())
}
