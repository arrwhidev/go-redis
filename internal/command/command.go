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

func CreateBulkString(s string) []byte {
	return []byte(fmt.Sprintf("%c%d\r\n%s\r\n", resp2.BulkStringByte, len(s), s))
}

func CreateNilBulkString() []byte {
	return []byte(fmt.Sprintf("%c-1\r\n", resp2.BulkStringByte))
}

func CreateArray(arr []string) []byte {
	res := []byte(fmt.Sprintf("%c%d\r\n", resp2.ArrayByte, len(arr)))
	for _, s := range arr {
		res = append(res, CreateBulkString(s)...)
	}
	return res
}

func CreateIntegerReply(i int) []byte {
	return []byte(fmt.Sprintf(":%d\r\n", i))
}
