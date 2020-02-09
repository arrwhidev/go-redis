package command

import "time"

type Clock interface {
	Now() time.Time
}

type realClock struct {}
func (realClock) Now() time.Time {
	return time.Now()
}