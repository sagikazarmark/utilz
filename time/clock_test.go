package time_test

import (
	"testing"

	stdtime "time"

	"github.com/sagikazarmark/utilz/time"
)

func TestStoppedClock_Now(t *testing.T) {
	ti := stdtime.Date(2017, stdtime.May, 10, 22, 52, 0, 0, stdtime.UTC)

	clock := time.NewStoppedClock(ti)

	if ti != clock.Now() {
		t.Errorf("expected clock's current time to be %v", ti)
	}
}

func TestClock_Now(t *testing.T) {
	ti := stdtime.Now()

	stdtime.Sleep(stdtime.Nanosecond)

	clock := time.NewClock()

	if ti = ti.Add(stdtime.Second); clock.Now().After(ti) {
		t.Errorf("expected clock's current time to be before %v", ti)
	}
}
