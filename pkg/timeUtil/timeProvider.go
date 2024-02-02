package timeUtil

import "time"

// TimeProvider defines an interface for getting the current time.
type TimeProvider interface {
	Now() time.Time
}

// DefaultTimeProvider returns the current time.
type DefaultTimeProvider struct{}

func (d *DefaultTimeProvider) Now() time.Time {
	return time.Now()
}
