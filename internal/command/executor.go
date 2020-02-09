package command

import (
	"errors"
	"fmt"
	"github.com/arrwhidev/go-redis/internal/store"
	"strconv"
	"strings"
	"time"
)

type Executor struct {
	Store *store.Store
	Clock
}

func NewExecutor(s *store.Store) *Executor {
	return &Executor{s, &realClock{}}
}

var commands = map[string]func(*Executor, []string) ([]byte, error){
	"ping": Ping,
	"echo": Echo,
	"quit": Quit,
	"set":  Set,
	"get":  Get,
}

func (e *Executor) Exec(cmd []string) []byte {
	head := strings.ToLower(cmd[0])
	fn := commands[head]
	if fn != nil {
		res, err := fn(e, cmd)
		if err != nil {
			return CreateError(err.Error())
		}
		return res
	}

	return CreateError(fmt.Sprintf("unknown command '%s'", cmd[0]))
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

// https://redis.io/commands/set
func Set(e *Executor, cmd []string) ([]byte, error) {
	size := len(cmd)
	if size < 3 {
		// Minimum command is `SET key value`
		return nil, errors.New("min args not met")
	}

	args, err := ToTuples(cmd[1:])
	if err != nil {
		return nil, err
	}

	var expiry int64 = -1

	for k, v := range args {
		if k == "EX" {
			seconds, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}

			now := e.Clock.Now()
			expiry = now.Add(time.Duration(seconds) * time.Second).UnixNano()
		} else if k == "PX" {
			ms, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}

			now := e.Clock.Now()
			expiry = now.Add(time.Duration(ms) * time.Millisecond).UnixNano()
		}
	}

	e.Store.Set(cmd[1], store.NewEntry(cmd[2], expiry))
	return CreateSimpleString("OK"), nil
}

// https://redis.io/commands/get
func Get(e *Executor, cmd []string) ([]byte, error) {
	v, err := e.Store.Get(cmd[1]) // TODO: handle array oob
	if err == nil {
		return CreateBulkString(v.Value), nil
	}
	return CreateNilBulkString(), nil
}

func ToTuples(arr []string) (map[string]string, error) {
	size := len(arr)
	if size%2 != 0 {
		return nil, errors.New("array length is not even")
	}

	m := make(map[string]string, size/2)
	for i := 0; i < size; i += 2 {
		m[arr[i]] = arr[i+1]
	}
	return m, nil
}
