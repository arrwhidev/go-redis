package command

import (
	"testing"
	"time"

	"github.com/arrwhidev/go-redis/internal/store"
)

type mockClock struct{}

func (mockClock) Now() time.Time {
	return time.Date(2100, time.January, 1, 1, 1, 1, 1, time.Local)
}

func mockNowAdd(d time.Duration) int64 {
	c := &mockClock{}
	return c.Now().Add(d).UnixNano()
}

func newExecutor() *Executor {
	return &Executor{store.NewStore(), &mockClock{}}
}

func TestItReturnsUnknown_whenUnknownCommand(t *testing.T) {
	res := newExecutor().Exec([]string{"WAT"})
	AssertError(t, "unknown command 'WAT'", res)
}
