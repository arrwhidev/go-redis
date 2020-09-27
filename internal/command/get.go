package command

import "errors"

// https://redis.io/commands/get
func Get(e *Executor, cmd []string) ([]byte, error) {
	size := len(cmd)
	if size < 1 {
		// Minimum command is `GET key`
		return nil, errors.New("min args not met")
	}

	return handleGet(e, cmd[0])
}

func handleGet(e *Executor, key string) ([]byte, error) {
	v, err := e.Store.Get(key)
	if err == nil {
		return CreateBulkString(v.Value), nil
	}
	return nil, errors.New(err.Error())
}
