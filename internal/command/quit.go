package command

func Quit(e *Executor, cmd []string) ([]byte, error) {
	return CreateSimpleString("OK"), nil
}
