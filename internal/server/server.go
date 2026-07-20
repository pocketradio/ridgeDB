package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"ridgeDB/internal/parser"
	"ridgeDB/internal/storage"
)

func Start() net.Listener {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	return listener
}

func Accept(db *storage.Store, listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go HandleConnection(db, conn)
	}
}

func HandleConnection(db *storage.Store, conn net.Conn) { // returns a *bufio.Reader
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {

		message, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Read error %v", err)
			return
		}

		parsed_message, err := parser.Parse(message)
		if err != nil {
			fmt.Println(err)
			continue
		}

		cmd, err := HandleCommand(parsed_message)

		if err != nil {
			fmt.Println(err)
			continue
		}

		ExecuteCommand(db, cmd)
	}

}
