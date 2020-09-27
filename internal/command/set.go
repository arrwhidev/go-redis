package command

import (
	"errors"
	"strconv"
	"time"

	"github.com/arrwhidev/go-redis/internal/store"
)

// https://redis.io/commands/set
func Set(e *Executor, cmd []string) ([]byte, error) {
	size := len(cmd)
	if size < 2 {
		// Minimum command is `SET key value`
		return nil, errors.New("min args not met")
	}

	args, err := ToTuples(cmd)
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

	e.Store.Set(cmd[0], store.NewEntry(cmd[1], expiry))
	return CreateSimpleString("OK"), nil
}
