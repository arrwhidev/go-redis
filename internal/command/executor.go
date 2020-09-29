package command

import (
	"errors"
	"fmt"
	"strings"

	"github.com/arrwhidev/go-redis/internal/store"
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
	"keys": Keys,
	"del":  Del,
}

func (e *Executor) Exec(cmd []string) []byte {
	head := strings.ToLower(cmd[0])
	fn := commands[head]
	if fn != nil {
		res, err := fn(e, cmd[1:])
		if err != nil {
			return CreateError(err.Error())
		}
		return res
	}

	return CreateError(fmt.Sprintf("unknown command '%s'", cmd[0]))
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
