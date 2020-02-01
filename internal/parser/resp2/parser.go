package resp2

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
)

// RESP2 parser
type Parser struct {
	reader *bufio.Reader
}

func NewParser(r *bufio.Reader) *Parser {
	return &Parser{
		reader: r,
	}
}

const (
	ArrayStart = '*'
	CR         = '\r'
	LF         = '\n'
	CRLF       = "\r\n"
)

// *1\r\n$7\r\nCOMMAND\r\n

func (p *Parser) Parse() (command string) {


		//
		//// Read first byte to detect parser types
		//b, err := r.Reader.ReadByte()
		//if err != nil {
		//	log.Fatal("Failed to read bytes", err)
		//	r.Connection.Close()
		//	return
		//}
		//
		//if b != ArrayStart {
		//	log.Fatal("Invalid protocol")
		//	r.Connection.Close()
		//	return
		//}
		//
		//// Read rest of the parser until LF.
		//fullCommand, err := r.Reader.ReadBytes('\n')
		//
		//// Trim CRLF
		//length := len(fullCommand)
		//trimmedCommand := fullCommand[:length-2]
		//
		//// Get length of parser.
		//commandLength, err := strconv.Atoi(string(trimmedCommand))
		//if err != nil {
		//	return
		//}
		//
		//if commandLength > 0 {
		//	fmt.Printf("%s\n",string(trimmedCommand))
		//}
		//
		//
		////buf := make([]byte, commandLength)
		////
		////value := string(buf)
		////fmt.Printf("%s\n", value)
		//
		//// clientAddr := c.RemoteAddr().String()
		//// message := string(data)
		//// clientAddr := c.RemoteAddr().String()
		//// fmt.Println(message + " from " + clientAddr + "\n")
		//// c.Write([]byte{b})

}

