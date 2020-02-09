package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
