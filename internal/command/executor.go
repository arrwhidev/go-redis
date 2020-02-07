package command

import (
	"fmt"
	"github.com/arrwhidev/go-redis/internal/store"
	"strconv"
	"strings"
	"time"
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
	size := len(cmd)
	if size < 3 {
		return CreateError("not enough parts"), nil
	}

	fmt.Println(cmd)

	var expiry int64 = -1
	if size > 3 {
		if cmd[3] == "EX" && size == 5 {
			seconds, err := strconv.Atoi(cmd[4])
			if err != nil {
				// TODO: handle
			}

			expiry = time.Now().Add(time.Duration(seconds) * time.Second).UnixNano()
		}
	}

	e.Store.Set(cmd[1], store.NewEntry(cmd[2], expiry))
	return CreateSimpleString("OK"), nil
}

func Get(e *Executor, cmd []string) ([]byte, error) {
	v, err := e.Store.Get(cmd[1]) // TODO: handle array oob
	if err == nil {
		return CreateBulkString(v.Value), nil
	}
	return CreateNilBulkString(), nil
}
