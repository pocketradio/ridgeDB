package server

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func HandleCommand(s []string) (Command, error) {

	method := strings.ToUpper(s[0])

	switch method {
	case "SET":
		if len(s) != 4 {
			return Command{}, fmt.Errorf("usage: SET <key> <value> <ttl_seconds>")
		}

		ttl, err := strconv.Atoi(s[3])
		if err != nil || ttl < 0 {
			return Command{}, fmt.Errorf("ttl must be a non-negative number")
		}

	case "GET":
		if len(s) != 2 {
			return Command{}, fmt.Errorf("usage: GET <key>")
		}

	case "DEL":
		if len(s) != 2 {
			return Command{}, fmt.Errorf("usage: DEL <key>")
		}

	default:
		return Command{}, fmt.Errorf("unknown command: %s", method)
	}

	if method == "SET" {
		ttl, _ := strconv.Atoi(s[3])
		cmd := Command{
			Data:   s[2],
			Method: method,
			Key:    s[1],
			TTL:    time.Duration(ttl) * time.Second,
		}
		return cmd, nil
	}

	cmd := Command{
		Data:   "",
		Method: method,
		Key:    s[1],
	}
	return cmd, nil

}
