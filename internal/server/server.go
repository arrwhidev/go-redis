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
//
//const (
//	ArrayStart = '*'
//	CR         = '\r'
//	LF         = '\n'
//	CRLF       = "\r\n"
//)
//
//func handleConnection(c net.Conn) {
//	for {
//		reader := bufio.NewReader(c)
//		b, err := reader.ReadByte()
//		if err != nil {
//			log.Fatal("Failed to read bytes", err)
//			c.Close()
//			return
//		}
//
//		if b != ArrayStart {
//			log.Fatal("Invalid protocol")
//			c.Close()
//			return
//		}
//
//		// Read rest of the parser
//		bytes, err := reader.ReadBytes('\n')
//
//		// Trim CRLF
//		length := len(bytes)
//		parser := bytes[:length-2]
//
//		commandLength, err := strconv.Atoi(string(parser))
//		if err != nil {
//			return
//		}
//
//		buf := make([]byte, strLen)
//
//		value := string(buf)
//		fmt.Printf("%s\n", value)
//
//		// clientAddr := c.RemoteAddr().String()
//		// message := string(data)
//		// clientAddr := c.RemoteAddr().String()
//		// fmt.Println(message + " from " + clientAddr + "\n")
//		// c.Write([]byte{b})
//	}
//}
