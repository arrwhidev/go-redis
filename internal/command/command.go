package command

import (
	"fmt"
	"github.com/arrwhidev/go-redis/internal/parser/resp2"
)

func CreateError(s string) []byte {
	return []byte(fmt.Sprintf("%cERR %s%s", resp2.ErrorByte, s, resp2.CRLF))
}

func CreateSimpleString(s string) []byte {
	return []byte(fmt.Sprintf("%c%s%s", resp2.StringByte, s, resp2.CRLF))
}
