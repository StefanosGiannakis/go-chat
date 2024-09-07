package main

import (
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/StefanosGiannakis/go-chat/server"
)

func main() {
	server := server.NewServer()
	http.Handle("/ws", websocket.Handler(server.HandleConn))
	http.ListenAndServe(":8080", nil)
}
