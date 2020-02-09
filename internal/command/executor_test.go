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
	res, _ := newExecutor().Exec([]string{"WAT"})
	assert.Equal(t, "-ERR unknown command 'WAT'\r\n", string(res))
}

func TestItReturnsPong_whenPingCommand(t *testing.T) {
	res, _ := newExecutor().Exec([]string{"PING"})
	assert.Equal(t, "+PONG\r\n", string(res))
}

func TestItReturnsMessage_whenEchoCommand(t *testing.T) {
	res, _ := newExecutor().Exec([]string{"ECHO", "Hello, world"})
	assert.Equal(t, "+Hello, world\r\n", string(res))
}

func TestItReturnsOK_whenQuitCommand(t *testing.T) {
	res, _ := newExecutor().Exec([]string{"QUIT"})
	assert.Equal(t, "+OK\r\n", string(res))
}

func TestGetReturnsBulkStringWithValue_whenKeyExists(t *testing.T) {
	e := newExecutor()
	e.Exec([]string{"SET", "hello", "world"})
	res, _ := e.Exec([]string{"GET", "hello"})
	assert.Equal(t, "$5\r\nworld\r\n", string(res))
}

func TestSetReturnsOK_whenSetWasSuccessful(t *testing.T) {
	res, _ := newExecutor().Exec([]string{"SET", "hello", "world"})
	assert.Equal(t, "+OK\r\n", string(res))
}

func TestSetReturnsError_whenNotEnoughParts(t *testing.T) {
	res, _ := newExecutor().Exec([]string{"SET", "hello"})
	assert.Equal(t, "-ERR not enough parts\r\n", string(res))
}

func TestSetWithEX(t *testing.T) {
	executor := newExecutor()
	res, _ := executor.Exec([]string{"SET", "hello", "world", "EX", "60"})
	assert.Equal(t, "+OK\r\n", string(res))

	e, _ := executor.Store.Get("hello")
	expectedTime := mockNowAdd(time.Duration(60) * time.Second)
	assert.Equal(t, expectedTime, e.Expires)
}
