package command

import (
	"testing"
)

func TestReturnsBulkStringWithValue_whenKeyExists(t *testing.T) {
	e := newExecutor()
	e.Exec([]string{"SET", "hello", "world"})
	res := e.Exec([]string{"GET", "hello"})
	AssertBulkString(t, "world", res)
}

func TestReturnsError_whenKeyNotExists(t *testing.T) {
	e := newExecutor()
	res := e.Exec([]string{"GET", "hello"})
	AssertError(t, "Key 'hello' not found", res)
}
