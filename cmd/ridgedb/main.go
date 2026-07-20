package main

import (
	"fmt"
	"ridgeDB/internal/server"
	"ridgeDB/internal/storage"
)

func main() {
	db := storage.NewStore()
	go db.StartCleanup()
	fmt.Println("Starting server on port 8000...")
	listener := server.Start()
	server.Accept(db, listener)
}
