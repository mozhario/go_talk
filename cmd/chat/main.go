package main

import (
	"fmt"
	"net/http"

	"github.com/mozhario/go_talk/chat/http_server"
	"github.com/mozhario/go_talk/chat/websocket"
	"github.com/mozhario/go_talk/config"
	"github.com/mozhario/go_talk/db"
)

func main() {
	fmt.Println("GoTalk v0.01")

	db.InitDB()
	defer db.CloseDB()

	go serveHTTP()
	go serveWebSocket()

	select {}
}

func serveHTTP() {
	fmt.Println("HTTP server startup")

	mux := http.NewServeMux()

	server := http_server.HTTPServer{
		Mux: mux,
	}
	server.Listen()
}

func serveWebSocket() {
	fmt.Println("Websocket server startup")

	pool := websocket.NewPool()

	server := websocket.WebSocketServer{
		Host: config.ServerHost,
		Port: config.WebscketPort,
		Pool: pool,
	}
	server.Listen()
}
