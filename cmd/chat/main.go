package main

import (
	"github.com/mozhario/go_talk/chat/server"
	"github.com/mozhario/go_talk/config"
)

func main() {
	server := server.WebSocketServer{
		Host: config.ServerHost,
		Port: config.ServerPort,
	}
	server.Listen()
}
