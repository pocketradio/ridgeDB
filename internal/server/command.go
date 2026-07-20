package server

import "ridgeDB/internal/storage"

type Command struct {
	Method string
	Key    string
	Data   string
	Expiry bool
}

type CommandResult struct {
	Status string // for set
	Value  storage.Value
	HasKey bool
}
