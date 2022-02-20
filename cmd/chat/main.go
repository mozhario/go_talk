package main

import (
	"github.com/mozhario/go_talk/chat/server"
	"github.com/mozhario/go_talk/settings"
)

func main() {
	server := server.Server{settings.ServerPort}
	server.Listen()
}
