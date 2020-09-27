package command

import "testing"

func TestItReturnsMessage_whenEchoCommand(t *testing.T) {
	res := newExecutor().Exec([]string{"ECHO", "Hello, world"})
	AssertSimpleString(t, "Hello, world", res)
}
