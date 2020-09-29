package command

import (
	"errors"
)

// https://redis.io/commands/del
func Del(e *Executor, cmd []string) ([]byte, error) {
	size := len(cmd)
	if size < 1 {
		// Minimum command is `DEL key`
		return nil, errors.New("min args not met")
	}

	num := 0
	for _, key := range cmd {
		// TODO: store could delete many keys under one mutex lock
		if e.Store.Del(key) {
			num++
		}
	}

	return CreateIntegerReply(num), nil
}
