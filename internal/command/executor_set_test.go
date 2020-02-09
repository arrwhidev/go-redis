package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestReturnsOK_whenSuccessful(t *testing.T) {
	res := newExecutor().Exec([]string{"SET", "hello", "world"})
	assert.Equal(t, "+OK\r\n", string(res))
}

func TestReturnsError_whenNotEnoughParts(t *testing.T) {
	res := newExecutor().Exec([]string{"SET", "hello"})
	assert.Equal(t, "-ERR min args not met\r\n", string(res))
}

func TestExpiresIsMinusOne_whenNotSpecified(t *testing.T) {
	executor := newExecutor()
	executor.Exec([]string{"SET", "hello", "world"})

	e, _ := executor.Store.Get("hello")
	assert.Equal(t, int64(-1), e.Expires)
}

func TestEXArgGivesExpiryOfNowPlusSeconds(t *testing.T) {
	executor := newExecutor()
	res := executor.Exec([]string{"SET", "hello", "world", "EX", "60"})
	assert.Equal(t, "+OK\r\n", string(res))

	e, _ := executor.Store.Get("hello")
	expectedTime := mockNowAdd(time.Duration(60) * time.Second)
	assert.Equal(t, expectedTime, e.Expires)
}

func TestSetWithPXArgGivesExpiryOfNowPlusMilliseconds(t *testing.T) {
	executor := newExecutor()
	res := executor.Exec([]string{"SET", "hello", "world", "PX", "1000"})
	assert.Equal(t, "+OK\r\n", string(res))

	e, _ := executor.Store.Get("hello")
	expectedTime := mockNowAdd(time.Duration(1000) * time.Millisecond)
	assert.Equal(t, expectedTime, e.Expires)
}

func TestLastArgOverwritesEarlier(t *testing.T) {
	executor := newExecutor()
	res := executor.Exec([]string{"SET", "hello", "world", "EX", "60", "PX", "1000"})
	assert.Equal(t, "+OK\r\n", string(res))

	e, _ := executor.Store.Get("hello")
	expectedTime := mockNowAdd(time.Duration(1000) * time.Millisecond)
	assert.Equal(t, expectedTime, e.Expires)
}

func TestReturnsError_whenOddArgs(t *testing.T) {
	res := newExecutor().Exec([]string{"SET", "hello", "world", "PX"})
	assert.Equal(t, "-ERR array length is not even\r\n", string(res))
}
