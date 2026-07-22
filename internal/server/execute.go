package server

import (
	"ridgeDB/internal/persistence"
	"ridgeDB/internal/storage"
	"time"
)

func ExecuteCommand(db *storage.Store, cmd Command, aof *persistence.AOF) (CommandResult, error) {
	switch cmd.Method {
	case "SET":

		err := aof.AppendSet(cmd.Key, cmd.Data)
		if err != nil {
			return CommandResult{}, err
		}
		if cmd.Expiry {
			_ = db.Set(cmd.Key, storage.Value{
				Data:      cmd.Data,
				HasExpiry: cmd.Expiry,
				ExpiresAt: time.Now().Add(240 * time.Hour),
			})

			return CommandResult{Status: "OK"}, nil
		}

		_ = db.Set(cmd.Key, storage.Value{
			Data:      cmd.Data,
			HasExpiry: cmd.Expiry,
			ExpiresAt: time.Time{},
		})

		return CommandResult{Status: "OK"}, nil

	case "GET":
		val, ok := db.Get(cmd.Key)
		if ok != true {
			return CommandResult{HasKey: false, Status: "Key not present."}, nil
		}

		return CommandResult{
			HasKey: true,
			Status: "Key present.",
			Value:  val,
		}, nil

	case "DEL":
		_ = db.Delete(cmd.Key)

		return CommandResult{
			Status: "Key deleted.",
		}, nil
	}

	return CommandResult{
		Status: "Please enter a valid method.",
	}, nil
}
