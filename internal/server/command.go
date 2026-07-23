package server

import (
	"ridgeDB/internal/storage"
	"time"
)

type Command struct {
	Method string
	Key    string
	Data   string
	TTL    time.Duration
}

type CommandResult struct {
	Status string // for set
	Value  storage.Value
	HasKey bool
}
