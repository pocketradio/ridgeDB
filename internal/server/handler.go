package server

import (
	"fmt"
	"strings"
)

func HandleCommand(s []string) (Command, error) {

	method := strings.ToUpper(s[0])

	switch method {
	case "SET":
		if len(s) != 4 {
			return Command{}, fmt.Errorf("usage: SET <key> <value> <expiry>")
		}

		expiry := strings.ToUpper(s[3])
		if expiry != "TRUE" && expiry != "FALSE" {
			return Command{}, fmt.Errorf("expiry must be true or false")
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
		cmd := Command{
			Data:   s[2],
			Method: method,
			Key:    s[1],
			Expiry: strings.ToUpper(s[3]) == "TRUE",
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
