package resp2

import (
	"bufio"
	"errors"
	"strconv"
)

// RESP2 parser
type Parser struct {
	reader *bufio.Reader
}

func NewParser(r *bufio.Reader) *Parser {
	return &Parser{r}
}

const (
	StringByte     = '+'
	ErrorByte      = '-'
	IntegerByte    = ':'
	BulkStringByte = '$'
	ArrayByte      = '*'
	CR             = '\r'
	LF             = '\n'
	CRLF           = "\r\n"
)

func (p *Parser) Parse() (command []string, err error) {
	b, err := p.reader.ReadByte()
	if err != nil {
		return nil, err
	}

	if b != ArrayByte {
		return nil, errors.New("invalid protocol")
	}
	numCommands, err := readNumCommands(p.reader)

	return make([]string, numCommands), nil
}

func readNumCommands(reader *bufio.Reader) (int, error) {
	// Read the rest of the bytes until LF.
	// looks like this; 3\r\n
	bytes, err := reader.ReadBytes(LF)
	if err != nil {
		return 0, err
	}

	// Strip off the CRLF at the end to get the num of commands.
	bytes = bytes[:len(bytes)-2]

	// Parse to int.
	numCommands, err := strconv.Atoi(string(bytes))
	if err != nil {
		return 0, err
	}

	return numCommands, nil
}
