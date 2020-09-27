package command

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReturnsError_whenNotEnoughArgs(t *testing.T) {
	res := newExecutor().Exec([]string{"KEYS"})
	AssertError(t, "min args not met", res)
}

func TestReturnsEmptyArray_whenNoKeys(t *testing.T) {
	res := newExecutor().Exec([]string{"KEYS", "*"})
	AssertEmptyArray(t, res)
}

func TestReturnsPopulatedArray_whenHasKeys(t *testing.T) {
	e := newExecutor()
	expected := "*1000\r\n"
	for i := 0; i < 1000; i++ {
		key := fmt.Sprintf("hello%d", i)
		e.Exec([]string{"SET", key, "world"})
		expected += fmt.Sprintf("$%d\r\n%s\r\n", len(key), key)
	}

	res := e.Exec([]string{"KEYS", "*"})
	assert.Len(t, string(res), len(expected)) // only asserting len because ordering
}
