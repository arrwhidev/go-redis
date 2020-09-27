package command

import "errors"

func Echo(e *Executor, cmd []string) ([]byte, error) {
	size := len(cmd)
	if size < 1 {
		// Minimum command is `ECHO message`
		return nil, errors.New("min args not met")
	}

	return CreateSimpleString(cmd[0]), nil
}
