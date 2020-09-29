package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateError(t *testing.T) {
	cmd := CreateError("unknown command 'WAT'")
	assert.Equal(t, "-ERR unknown command 'WAT'\r\n", string(cmd))
}

func TestCreateSimpleString(t *testing.T) {
	cmd := CreateSimpleString("PONG")
	assert.Equal(t, "+PONG\r\n", string(cmd))
}

func TestCreateBulkString(t *testing.T) {
	cmd := CreateBulkString("TESTING")
	assert.Equal(t, "$7\r\nTESTING\r\n", string(cmd))
}

func TestCreateNilBulkString(t *testing.T) {
	cmd := CreateNilBulkString()
	assert.Equal(t, "$-1\r\n", string(cmd))
}

func TestCreateArray_whenEmpty(t *testing.T) {
	cmd := CreateArray([]string{})
	assert.Equal(t, "*0\r\n", string(cmd))
}

func TestCreateArray_whenNotEmpty(t *testing.T) {
	cmd := CreateArray([]string{"hi", "world"})
	assert.Equal(t, "*2\r\n$2\r\nhi\r\n$5\r\nworld\r\n", string(cmd))
}

func TestCreateIntegerReply(t *testing.T) {
	cmd := CreateIntegerReply(739)
	assert.Equal(t, ":739\r\n", string(cmd))
}
