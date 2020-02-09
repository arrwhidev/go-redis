package command

import (
	"github.com/arrwhidev/go-redis/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockClock struct{}

func (mockClock) Now() time.Time {
	return time.Date(2100, time.January, 1, 1, 1, 1, 1, time.Local)
}

func mockNowAdd(d time.Duration) int64 {
	c := &mockClock{}
	return c.Now().Add(d).UnixNano()
}

//

func newExecutor() *Executor {
	return &Executor{store.NewStore(), &mockClock{}}
}

func TestItReturnsUnknown_whenUnknownCommand(t *testing.T) {
	res := newExecutor().Exec([]string{"WAT"})
	assert.Equal(t, "-ERR unknown command 'WAT'\r\n", string(res))
}

func TestItReturnsPong_whenPingCommand(t *testing.T) {
	res := newExecutor().Exec([]string{"PING"})
	assert.Equal(t, "+PONG\r\n", string(res))
}

func TestItReturnsMessage_whenEchoCommand(t *testing.T) {
	res := newExecutor().Exec([]string{"ECHO", "Hello, world"})
	assert.Equal(t, "+Hello, world\r\n", string(res))
}

func TestItReturnsOK_whenQuitCommand(t *testing.T) {
	res := newExecutor().Exec([]string{"QUIT"})
	assert.Equal(t, "+OK\r\n", string(res))
}
