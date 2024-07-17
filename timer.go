package uhelpers

import "time"

type timer struct {
	startTime time.Time
}

func NewTimer() timer {
	return timer{startTime: time.Now()}
}

func (t *timer) Start() {
	t.startTime = time.Now()
}

func (t *timer) MeasureAndReset() time.Duration {
	ret := time.Since(t.startTime)
	t.startTime = time.Now()
	return ret
}
