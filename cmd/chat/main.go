package main

import (
	"fmt"
	"net/http"

	"github.com/mozhario/go_talk/routes"

	"github.com/mozhario/go_talk/chat/server"
	"github.com/mozhario/go_talk/config"
)

func main() {
	fmt.Println("GoTalk v0.01")

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

	server := server.WebSocketServer{
		Host: config.ServerHost,
		Port: config.WebscketPort,
	}
	server.Listen()
}
