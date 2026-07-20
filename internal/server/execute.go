package server

import (
	"ridgeDB/internal/storage"
	"time"
)

type CommandResult struct {
	Status string // for set
	Value  storage.Value
	HasKey bool
}

func ExecuteCommand(db *storage.Store, cmd Command) CommandResult {
	switch cmd.Method {
	case "SET":
		_ = db.Set(cmd.Key, storage.Value{
			Data:      cmd.Data,
			HasExpiry: true,
			ExpiresAt: time.Now().Add(240 * time.Hour),
		})

		return CommandResult{Status: "OK"}

	case "GET":
		val, ok := db.Get(cmd.Key)
		if ok != true {
			return CommandResult{HasKey: false, Status: "Key not present."}
		}

		return CommandResult{
			HasKey: true,
			Status: "Key present.",
			Value:  val,
		}

	case "DEL":
		_ = db.Delete(cmd.Key)

		return CommandResult{
			Status: "Key deleted.",
		}
	}

	return CommandResult{
		Status: "Please enter a valid method.",
	}
}
