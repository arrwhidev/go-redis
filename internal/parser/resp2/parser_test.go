package resp2

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func newParser(in string) *Parser {
	reader := bufio.NewReader(strings.NewReader(in))
	return NewParser(reader)
}

func TestItReturnsError_whenInvalidProtocol(t *testing.T) {
	_, err := newParser("wat").Parse()
	assert.NotNil(t, err)
	assert.Errorf(t, err, "invalid protocol")
}

func TestItReturnsEmptyArray_whenEmptyArray(t *testing.T) {
	cmd, _ := newParser("*0\r\n").Parse()
	assert.Empty(t, cmd, cmd)
}

func TestItReturnsArrayWithOne_whenArrayHasOne(t *testing.T) {
	cmd, _ := newParser("*1\r\n$7\r\nCOMMAND\r\n").Parse()
	assert.Len(t, cmd, 1)
}

func TestItReturnsArrayWithCommand_whenArrayHasOne(t *testing.T) {
	cmd, _ := newParser("*1\r\n$7\r\nCOMMAND\r\n").Parse()
	assert.Equal(t, "COMMAND", cmd[0])
}

func TestItReturnsArrayWithCommandAndArgs_whenArrayHasMoreThanOne(t *testing.T) {
	cmd, _ := newParser("*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n").Parse()
	assert.Len(t, cmd, 3)
	assert.ElementsMatch(t, cmd, []string{ "SET", "key", "value" })
}
