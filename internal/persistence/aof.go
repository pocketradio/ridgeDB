package persistence

import (
	"fmt"
	"os"
	"sync"
)

type AOF struct {
	File     *os.File
	Filepath string
	mu       sync.Mutex
}

func NewAOF(path string) (*AOF, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}

	return &AOF{File: f, Filepath: path}, nil
}

func Open() (*AOF, error) {
	return NewAOF("ridgedb.aof")
}

func (a *AOF) AppendSet(key, value string, expiry bool) error {

	a.mu.Lock()
	defer a.mu.Unlock()

	_, err := fmt.Fprintf(a.File, "SET %s %s %t\n", key, value, expiry)

	if err != nil {
		return err
	}

	if err := a.File.Sync(); err != nil {
		return err
	}

	return nil
}

func (a *AOF) AppendDel(key string) error {

	a.mu.Lock()
	defer a.mu.Unlock()

	_, err := fmt.Fprintf(a.File, "DEL %s\n", key)

	if err != nil {
		return err
	}

	if err := a.File.Sync(); err != nil {
		return err
	}

	return nil
}
