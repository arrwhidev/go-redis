package command

import "errors"

// https://redis.io/commands/keys
func Keys(e *Executor, cmd []string) ([]byte, error) {
	size := len(cmd)
	if size < 1 {
		// Minimum command is `KEYS pattern`
		return nil, errors.New("min args not met")
	}

	return handleKeys(e, cmd[0])
}

func handleKeys(e *Executor, pattern string) ([]byte, error) {
	// TODO: use pattern for proper keys search
	return CreateArray(e.Store.Keys()), nil
}
