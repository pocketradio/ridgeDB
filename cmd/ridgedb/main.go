package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"ridgeDB/internal/persistence"
	"ridgeDB/internal/server"
	"ridgeDB/internal/storage"
	"strconv"
)

func main() {

	portPtr := flag.String("port", ":8000", "TCP port")

	flag.Parse()

	_, port, err := net.SplitHostPort(*portPtr)
	if err != nil {
		log.Fatalf("invalid port %q: expected format like :8000", *portPtr)
	}

	if _, err := strconv.Atoi(port); err != nil {
		log.Fatalf("invalid port %q: port must be a number", *portPtr)
	}

	fmt.Println("Port: ", *portPtr)

	db := storage.NewStore()

	aof, err := persistence.Open()
	if err != nil {
		log.Fatalf("failed to open AOF: %v", err)
	}

	listener := server.Start(*portPtr)
	go db.StartCleanup()

	srv := server.NewServer(db, aof, listener)
	srv.Accept()

}
