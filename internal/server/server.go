package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"ridgeDB/internal/parser"
	"ridgeDB/internal/storage"
)

func Start(port string) net.Listener {

	listener, err := net.Listen("tcp", port)
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

func HandleConnection(db *storage.Store, conn net.Conn) {
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
			fmt.Fprintln(conn, err)
			continue
		}

		cmd, err := HandleCommand(parsed_message)

		if err != nil {
			fmt.Fprintln(conn, err)
			continue
		}

		result := ExecuteCommand(db, cmd)

		switch cmd.Method {
		case "SET":
			fmt.Fprintln(conn, result.Status)

		case "GET":
			if !result.HasKey {
				fmt.Fprintln(conn, result.Status)
			} else {
				fmt.Fprintln(conn, result.Value.Data)
			}

		case "DEL":
			fmt.Fprintln(conn, result.Status)

		default:
			fmt.Fprintln(conn, result.Status)
		}

	}

}
