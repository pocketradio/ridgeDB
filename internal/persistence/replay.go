package persistence

import (
	"bufio"
	"fmt"
	"os"
	"ridgeDB/internal/storage"
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

		expiry := strings.ToUpper(parsedLine[3]) == "TRUE"
		value := storage.Value{
			Data:      parsedLine[2],
			HasExpiry: expiry,
		}

		if expiry {
			value.ExpiresAt = time.Now().Add(240 * time.Hour)
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
