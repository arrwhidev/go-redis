package command

import "testing"

func TestItReturnsOK_whenQuitCommand(t *testing.T) {
	res := newExecutor().Exec([]string{"QUIT"})
	AssertSimpleString(t, "OK", res)
}
