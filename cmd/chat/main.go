package main

import (
	"fmt"
	"net/http"

	"github.com/mozhario/go_talk/routes"

	"github.com/mozhario/go_talk/chat/db"
	"github.com/mozhario/go_talk/chat/websocket"
	"github.com/mozhario/go_talk/config"
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
	routes.SetupRoutes(mux)
	http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), mux)
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
