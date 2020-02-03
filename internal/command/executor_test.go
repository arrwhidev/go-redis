package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func newExecutor() *Executor {
	return NewExecutor()
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


