package protocol

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func Decode(r *bufio.Reader) ([]string, error) {
	marker, err := r.ReadByte()
	if err != nil {
		return []string{}, err
	}

	if marker != '*' {
		return nil, fmt.Errorf("expected array")
	}

	length, err := r.ReadString('\n')
	if err != nil {
		return nil, err
	}

	n, err := strconv.Atoi(strings.TrimSpace(length))

	if err != nil {
		return nil, err
	}

	if n <= 0 {
		return nil, fmt.Errorf("invalid array length")
	}

	// n will be the number of arguments the command will contain. for instance set name x -> n = 3

	// eg stream : *4\r\n$3\r\nSET\r\n$3\r\nfoo\r\n$5\r\nhello\r\n$2\r\n10\r\n

	args := make([]string, 0, n)

	for i := 0; i < n; i++ {
		marker, err = r.ReadByte()
		if err != nil {
			return nil, err
		}

		if marker != '$' {
			return nil, fmt.Errorf("invalid command")
		}

		bulk, err := r.ReadString('\n')
		if err != nil {
			return nil, err
		}

		bulkLen, err := strconv.Atoi(strings.TrimSpace(bulk))

		if err != nil {
			return nil, err
		}

		if bulkLen < 0 {
			return nil, fmt.Errorf("invalid bulk length.")
		}

		buffer := make([]byte, bulkLen)

		_, err = io.ReadFull(r, buffer)
		if err != nil {
			return nil, err
		}

		crlf := make([]byte, 2)
		_, err = io.ReadFull(r, crlf)
		if err != nil {
			return nil, err
		}
		if string(crlf) != "\r\n" {
			return nil, fmt.Errorf("expected CRLF")
		}

		args = append(args, string(buffer))
	}

	return args, nil
}
