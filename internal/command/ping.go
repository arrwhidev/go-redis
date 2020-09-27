package command

func Ping(e *Executor, cmd []string) ([]byte, error) {
	return CreateSimpleString("PONG"), nil
}
