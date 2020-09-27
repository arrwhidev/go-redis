package command

import (
	"testing"
)

func TestPingReturnsPong(t *testing.T) {
	e := newExecutor()
	res, _ := Ping(e, []string{})
	AssertSimpleString(t, "PONG", res)
}
