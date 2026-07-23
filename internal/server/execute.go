package server

import (
	"ridgeDB/internal/persistence"
	"ridgeDB/internal/storage"
	"time"
)

func ExecuteCommand(db *storage.Store, cmd Command, aof *persistence.AOF) (CommandResult, error) {
	switch cmd.Method {
	case "SET":
		value := storage.Value{
			Data: cmd.Data,
		}

		if cmd.TTL > 0 {
			value.HasExpiry = true
			value.ExpiresAt = time.Now().Add(cmd.TTL)
		}

		err := aof.AppendSet(cmd.Key, cmd.Data, value.ExpiresAt)
		if err != nil {
			return CommandResult{}, err
		}

		_ = db.Set(cmd.Key, value)
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
		err := aof.AppendDel(cmd.Key)
		if err != nil {
			return CommandResult{}, err
		}

		_ = db.Delete(cmd.Key)

		return CommandResult{
			Status: "Key deleted.",
		}, nil
	}

	return CommandResult{
		Status: "Please enter a valid method.",
	}, nil
}
