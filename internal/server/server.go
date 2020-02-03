package server

import (
	"github.com/arrwhidev/go-redis/internal/request"
	"log"
	"net"
)

// Start a tcp server
func Start() {
	l, err := net.Listen("tcp4", ":6379")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal("Failed to accept", err)
			return
		}

		r := request.NewRequest(c)
		go r.Handle()
	}
}
