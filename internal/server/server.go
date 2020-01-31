package server

import (
	"bufio"
	"fmt"
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
		go handleConnection(c)
	}
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		data, err := bufio.NewReader(c).ReadBytes('\n')
		if err != nil {
			log.Fatal("Failed to read bytes", err)
			c.Close()
			return
		}

		message := string(data)
		clientAddr := c.RemoteAddr().String()
		fmt.Println(message + " from " + clientAddr + "\n")
		c.Write([]byte("+PONG"))
	}
}
