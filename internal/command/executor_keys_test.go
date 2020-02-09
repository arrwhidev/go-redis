package command

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsError_whenNotEnoughArgs(t *testing.T) {
	res := newExecutor().Exec([]string{"KEYS"})
	assert.Equal(t, "-ERR min args not met\r\n", string(res))
}

func TestReturnsEmptyArray_whenNoKeys(t *testing.T) {
	res := newExecutor().Exec([]string{"KEYS", "*"})
	assert.Equal(t, "*0\r\n", string(res))
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
