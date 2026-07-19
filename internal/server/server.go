package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"ridgeDB/internal/parser"
)

func Start() net.Listener {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	return listener
}

func Accept(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go HandleConnection(conn)
	}
}

func HandleConnection(conn net.Conn) { // returns a *bufio.Reader
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

		err = HandleCommand(parsed_message)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

}
