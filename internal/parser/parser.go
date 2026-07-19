package parser

import (
	"fmt"
	"strings"
)

func Parse(s string) ([]string, error) {
	message := strings.Fields(s)

	if len(message) == 0 {
		return nil, fmt.Errorf("empty command")
	}

	return message, nil
}
