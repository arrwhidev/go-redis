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
}

func (e *Executor) Exec(cmd []string) ([]byte, error) {
	head := strings.ToLower(cmd[0])
	fn := commands[head]
	if fn != nil {
		return fn(cmd)
	}

	return CreateError(fmt.Sprintf("unknown command '%s'", cmd[0])), nil
}

func Ping(command []string) ([]byte, error) {
	return CreateSimpleString("PONG"), nil
}
