package server

import (
	"fmt"
	"io"

	"golang.org/x/net/websocket"
)

type Server struct {
	conns map[*websocket.Conn]bool
}

func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleConn(ws *websocket.Conn) {
	fmt.Println("New connection from ", ws.RemoteAddr())

	s.conns[ws] = true

	fmt.Println(s.conns)
	s.readLoop(ws)
}

func (s *Server) readLoop(ws *websocket.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := ws.Read(buf)

		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection closed by client")
				return
			}
			fmt.Println("Error reading from connection: ", err)
			continue
		}
		msg := buf[:n]

		s.broadcastMessage(msg, ws)
	}
}

func (s *Server) broadcastMessage(b []byte, exclCon *websocket.Conn) {
	for ws := range s.conns {
		if ws == exclCon {
			continue
		}
		go func(ws *websocket.Conn) {
			if _, err := ws.Write(b); err != nil {
				fmt.Println("Error writing to connection: ", err)
			}
		}(ws)
	}
}
