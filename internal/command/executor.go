package command

import (
	"fmt"
	"strings"
)

type Executor struct{}

func NewExecutor() *Executor {
	return &Executor{}
}

var commands = map[string]func([]string) ([]byte, error){
	"ping": Ping,
	"echo": Echo,
}

func (e *Executor) Exec(cmd []string) ([]byte, error) {
	head := strings.ToLower(cmd[0])
	fn := commands[head]
	if fn != nil {
		return fn(cmd)
	}

	return CreateError(fmt.Sprintf("unknown command '%s'", cmd[0])), nil
}

func Ping(cmd []string) ([]byte, error) {
	return CreateSimpleString("PONG"), nil
}

func Echo(cmd []string) ([]byte, error) {
	return CreateSimpleString(cmd[1]), nil
}
