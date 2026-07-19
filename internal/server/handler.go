package server

import (
	"fmt"
	"strings"
)

func HandleCommand(s []string) error {
	method := strings.ToUpper(s[0])

	switch method {
	case "SET":
		if len(s) != 3 {
			return fmt.Errorf("usage: SET <key> <value>")
		}

		userKey := s[1]
		value := s[2]

		_ = userKey
		_ = value

	case "GET":
		if len(s) != 2 {
			return fmt.Errorf("usage: GET <key>")
		}

		userKey := s[1]

		_ = userKey

	case "DEL":
		if len(s) != 2 {
			return fmt.Errorf("usage: DEL <key>")
		}

		userKey := s[1]

		_ = userKey

	default:
		return fmt.Errorf("unknown command: %s", method)
	}

	return nil
}
