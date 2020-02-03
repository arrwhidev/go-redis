package request

import (
	"bufio"
	"fmt"
	"github.com/arrwhidev/go-redis/internal/command"
	"github.com/arrwhidev/go-redis/internal/parser/resp2"
	"net"
)

type Request struct {
	Connection net.Conn
	Reader     *bufio.Reader
}

func NewRequest(c net.Conn) *Request {
	return &Request{
		Connection: c,
		Reader:     bufio.NewReader(c),
	}
}

func (r *Request) Handle() {
	p := resp2.NewParser(r.Reader)
	for {
		cmd, err := p.Parse()
		if err != nil {
			fmt.Println("failed to parse command", err)
			return
		}

		res, err := command.NewExecutor().Exec(cmd)
		if err != nil {
			fmt.Println("failed to execute command")
			return
		}

		r.Connection.Write(res)
	}
}
