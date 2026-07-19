package main

import (
	"log"
	"net"
	"ridgeDB/internal/server"
	"ridgeDB/internal/storage"
)

func main() {
	db := storage.NewStore()
	go db.StartCleanup()

	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go server.HandleConnection(conn)
	}
}
