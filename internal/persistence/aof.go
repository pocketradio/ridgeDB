package persistence

import (
	"fmt"
	"os"
)

type AOF struct {
	file *os.File
}

func NewAOF(filePath string) (*AOF, error) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	return &AOF{file: f}, nil
}

func Open() (*AOF, error) {
	return NewAOF("ridgedb.aof")
}

func (a *AOF) AppendSet(key, value string) error {

	_, err := fmt.Fprintf(a.file, "SET %s %s\n", key, value)

	if err != nil {
		return err
	}

	if err := a.file.Sync(); err != nil {
		return err
	}

	return nil
}
