package command

import (
	"testing"
)

func TestDelReturnsError_whenNotEnoughParts(t *testing.T) {
	res := newExecutor().Exec([]string{"DEL"})
	AssertError(t, "min args not met", res)
}

func TestDelReturnsIntegerReplyWithNumberOfKeysDeleted(t *testing.T) {
	e := newExecutor()
	e.Exec([]string{"SET", "key1", "hello"})
	e.Exec([]string{"SET", "key2", "world"})

	res := e.Exec([]string{"DEL", "key1", "key2", "key3"})
	AssertIntegerReply(t, 2, res)
}

func TestDelActuallyDeletes(t *testing.T) {
	e := newExecutor()
	e.Exec([]string{"SET", "key1", "hello"})
	e.Exec([]string{"DEL", "key1"})
	res := e.Exec([]string{"GET", "key1"})
	AssertError(t, "Key 'key1' not found", res)
}
