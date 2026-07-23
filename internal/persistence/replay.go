package persistence

import (
	"bufio"
	"fmt"
	"os"
	"ridgeDB/internal/storage"
	"strconv"
	"strings"
	"time"
)

func (a *AOF) Replay(db *storage.Store) error {
	f, err := os.Open(a.Filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if err := ReplayParse(db, scanner.Text()); err != nil {
			return err
		}
	}

	return scanner.Err()
}

func ReplayParse(db *storage.Store, s string) error {
	parsedLine := strings.Fields(s)
	if len(parsedLine) == 0 {
		return nil
	}

	switch strings.ToUpper(parsedLine[0]) {
	case "SET":
		if len(parsedLine) != 4 {
			return fmt.Errorf("invalid SET in AOF")
		}

		expiryUnix, err := strconv.ParseInt(parsedLine[3], 10, 64)
		if err != nil || expiryUnix < 0 {
			return fmt.Errorf("invalid expiry in AOF")
		}

		value := storage.Value{
			Data:      parsedLine[2],
			HasExpiry: expiryUnix > 0,
		}

		if expiryUnix > 0 {
			value.ExpiresAt = time.Unix(expiryUnix, 0)
			if value.ExpiresAt.Before(time.Now()) {
				return nil
			}
		}

		db.Set(parsedLine[1], value)

	case "DEL":
		if len(parsedLine) != 2 {
			return fmt.Errorf("invalid DEL in AOF")
		}

		db.Delete(parsedLine[1])

	default:
		return fmt.Errorf("invalid command in AOF: %s", parsedLine[0])
	}

	return nil
}
