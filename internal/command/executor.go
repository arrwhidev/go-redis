package command

import (
	"fmt"
	"github.com/arrwhidev/go-redis/internal/store"
	"strings"
)

type Executor struct {
	Store *store.Store
}

func NewExecutor(s *store.Store) *Executor {
	return &Executor{s}
}

var commands = map[string]func(*Executor, []string) ([]byte, error){
	"ping": Ping,
	"echo": Echo,
	"quit": Quit,
	"set":  Set,
	"get":  Get,
}

func (e *Executor) Exec(cmd []string) ([]byte, error) {
	head := strings.ToLower(cmd[0])
	fn := commands[head]
	if fn != nil {
		return fn(e, cmd)
	}

	return CreateError(fmt.Sprintf("unknown command '%s'", cmd[0])), nil
}

func Ping(e *Executor, cmd []string) ([]byte, error) {
	return CreateSimpleString("PONG"), nil
}

func Echo(e *Executor, cmd []string) ([]byte, error) {
	return CreateSimpleString(cmd[1]), nil // TODO: handle array oob
}

func Quit(e *Executor, cmd []string) ([]byte, error) {
	return CreateSimpleString("OK"), nil
}

func Set(e *Executor, cmd []string) ([]byte, error) {
	e.Store.Set(cmd[1], cmd[2]) // TODO: handle array oob
	return CreateSimpleString("OK"), nil
}

func Get(e *Executor, cmd []string) ([]byte, error) {
	v, err := e.Store.Get(cmd[1]) // TODO: handle array oob
	if err == nil {
		return CreateBulkString(v), nil
	}
	return CreateNilBulkString(), nil
}
