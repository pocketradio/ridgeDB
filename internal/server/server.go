package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"ridgeDB/internal/parser"
	"ridgeDB/internal/persistence"
	"ridgeDB/internal/storage"
)

type Server struct {
	DB       *storage.Store
	AOF      *persistence.AOF
	Listener net.Listener
}

func Start(port string) net.Listener {

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	return listener
}

func NewServer(db *storage.Store, aof *persistence.AOF, l net.Listener) *Server {
	return &Server{
		DB:       db,
		AOF:      aof,
		Listener: l,
	}
}

func (s *Server) Accept() {
	for {
		conn, err := s.Listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go s.HandleConnection(conn)
	}
}

func (s *Server) HandleConnection(conn net.Conn) {
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

		result := ExecuteCommand(s.DB, cmd)

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
