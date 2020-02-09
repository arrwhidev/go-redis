package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturnsBulkStringWithValue_whenKeyExists(t *testing.T) {
	e := newExecutor()
	e.Exec([]string{"SET", "hello", "world"})
	res := e.Exec([]string{"GET", "hello"})
	assert.Equal(t, "$5\r\nworld\r\n", string(res))
}
