package command

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReturnsError_whenNotEnoughParts(t *testing.T) {
	res := newExecutor().Exec([]string{"SET", "hello"})
	AssertError(t, "min args not met", res)
}

func TestReturnsError_whenOddArgs(t *testing.T) {
	res := newExecutor().Exec([]string{"SET", "hello", "world", "PX"})
	AssertError(t, "array length is not even", res)
}

func TestReturnsOK_whenSettingExistingKey(t *testing.T) {
	e := newExecutor()
	e.Exec([]string{"SET", "hello", "world"})
	res := e.Exec([]string{"SET", "hello", "world2"})
	AssertSimpleString(t, "OK", res)

	v, _ := e.Store.Get("hello")
	assert.Equal(t, "world2", v.Value)
}

func TestReturnsOK_whenSuccessful(t *testing.T) {
	res := newExecutor().Exec([]string{"SET", "hello", "world"})
	AssertSimpleString(t, "OK", res)
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
	AssertSimpleString(t, "OK", res)

	e, _ := executor.Store.Get("hello")
	expectedTime := mockNowAdd(time.Duration(60) * time.Second)
	assert.Equal(t, expectedTime, e.Expires)
}

func TestSetWithPXArgGivesExpiryOfNowPlusMilliseconds(t *testing.T) {
	executor := newExecutor()
	res := executor.Exec([]string{"SET", "hello", "world", "PX", "1000"})
	AssertSimpleString(t, "OK", res)

	e, _ := executor.Store.Get("hello")
	expectedTime := mockNowAdd(time.Duration(1000) * time.Millisecond)
	assert.Equal(t, expectedTime, e.Expires)
}

func TestLastArgOverwritesEarlier(t *testing.T) {
	executor := newExecutor()
	res := executor.Exec([]string{"SET", "hello", "world", "EX", "60", "PX", "1000"})
	AssertSimpleString(t, "OK", res)

	e, _ := executor.Store.Get("hello")
	expectedTime := mockNowAdd(time.Duration(1000) * time.Millisecond)
	assert.Equal(t, expectedTime, e.Expires)
}
