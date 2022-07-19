package request

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"github.com/arrwhidev/go-redis/internal/command"
	"github.com/arrwhidev/go-redis/internal/parser/resp2"
	"github.com/arrwhidev/go-redis/internal/store"
)

type Request struct {
	Connection net.Conn
	Reader     *bufio.Reader
	Store      *store.Store
}

func NewRequest(c net.Conn) *Request {
	return &Request{
		Connection: c,
		Reader:     bufio.NewReader(c),
		Store:      store.Instance(),
	}
}

func (r *Request) Handle() {
	p := resp2.NewParser(r.Reader)
	for {
		cmd, err := p.Parse()
		if err == io.EOF {
			return
		}

		if err != nil {
			fmt.Println("failed to parse command", err)
			return
		}

		res := command.NewExecutor(r.Store).Exec(cmd)
		r.Connection.Write(res)
	}
}
