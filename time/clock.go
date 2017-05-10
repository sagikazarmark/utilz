package time

import "time"

// Clock tells the time.
type Clock interface {
	// Now tells the actual time.
	Now() time.Time
}

// StoppedClock shows the moment it has been stopped.
type StoppedClock struct {
	t time.Time
}

// NewStoppedClock returns a new StoppedClock.
func NewStoppedClock(t time.Time) Clock {
	return &StoppedClock{t}
}

// Now tells the time when it has been stopped.
func (c *StoppedClock) Now() time.Time {
	return c.t
}

// RunningClock shows the current time.
type RunningClock struct{}

// NewClock returns a new running clock.
func NewClock() Clock {
	return &RunningClock{}
}

// Now tells the current time.
func (c *RunningClock) Now() time.Time {
	return time.Now()
}
