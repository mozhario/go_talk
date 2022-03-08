package main

import (
	"github.com/mozhario/go_talk/chat/server"
	"github.com/mozhario/go_talk/config"
)

func main() {
	server := server.WebSocketServer{config.ServerPort}
	server.Listen()
}
