package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func AssertSimpleString(t *testing.T, expected string, actual []byte) {
	assert.Equal(t, string(CreateSimpleString(expected)), string(actual))
}

func AssertBulkString(t *testing.T, expected string, actual []byte) {
	assert.Equal(t, string(CreateBulkString(expected)), string(actual))
}

func AssertError(t *testing.T, expected string, actual []byte) {
	assert.Equal(t, string(CreateError(expected)), string(actual))
}

func AssertArray(t *testing.T, expected []string, actual []byte) {
	assert.Equal(t, string(CreateArray(expected)), string(actual))
}

func AssertEmptyArray(t *testing.T, actual []byte) {
	AssertArray(t, []string{}, actual)
}
