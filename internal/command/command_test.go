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
