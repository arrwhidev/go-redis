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

	numCommands, err := readInteger(p.reader)
	commands := make([]string, numCommands)
	for i := 0; i < numCommands; i++ {
		b, err := p.reader.ReadByte()
		if err != nil {
			return nil, err
		}

		if b == BulkStringByte {
			bulkStr, err := readBulkString(p.reader)
			if err != nil {
				return nil, err
			}
			commands[i] = bulkStr
		} else {
			panic("unsupported") // TODO: temp panic whilst building
		}
	}

	return commands, nil
}

// Read bulk string.
func readBulkString(reader *bufio.Reader) (string, error) {
	// Read length of the bulk string.
	strLen, err := readInteger(reader)
	if err != nil {
		return "", err
	}

	// Read the bulk string.
	buf := make([]byte, strLen)
	_, err = reader.Read(buf)
	if err != nil {
		return "", err
	}

	// Discard the CRLF at the end.
	reader.Discard(2)

	return string(buf), nil
}

// Read bytes until CRLF and treat as an integer.
func readInteger(reader *bufio.Reader) (int, error) {
	// Read the rest of the bytes until LF.
	// looks like this; 3\r\n
	bytes, err := reader.ReadBytes(LF)
	if err != nil {
		return 0, err
	}

	bytes = stripCRLF(bytes)

	// Parse to int.
	n, err := strconv.Atoi(string(bytes))
	if err != nil {
		return 0, err
	}

	return n, nil
}

// Strip CRLF from end; 'HELLO\r\n' -> 'HELLO'.
func stripCRLF(bytes []byte) []byte {
	return bytes[:len(bytes)-2]
}
