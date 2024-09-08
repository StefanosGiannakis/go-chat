package main

import (
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/StefanosGiannakis/go-chat/server"
)

func main() {
	server := server.NewServer()
	http.Handle("/ws", websocket.Handler(server.HandleConn))

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)

}
