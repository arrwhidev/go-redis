package command

import (
	"github.com/arrwhidev/go-redis/internal/store"
	"github.com/stretchr/testify/assert"
	"testing"
)

func newExecutor() *Executor {
	return NewExecutor(store.NewStore())
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

func TestSetReturnsOK_whenSetWasSuccessful(t *testing.T) {
	res, _ := newExecutor().Exec([]string{"SET", "hello", "world"})
	assert.Equal(t, "+OK\r\n", string(res))
}

func TestGetReturnsBulkStringWithValue_whenKeyExists(t *testing.T) {
	e := newExecutor()
	e.Exec([]string{"SET", "hello", "world"})
	res, _ := e.Exec([]string{"GET", "hello"})
	assert.Equal(t, "$5\r\nworld\r\n", string(res))
}
