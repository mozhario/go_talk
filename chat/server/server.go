package server

import (
	"fmt"
	"log"
	"net"

	"github.com/mozhario/go_talk/chat/constants"
	"github.com/mozhario/go_talk/config"
)

type WebSocketServer struct {
	Host string
	Port string
}

func (server WebSocketServer) Listen() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.Port))
	if err != nil {
		log.Fatal(constants.ErrorListenerFailedToStart + err.Error())
	}

	for {
		conn, _ := listener.Accept()
		buff := make([]byte, config.ServerBufferSize)
		msg, err := conn.Read(buff[0:])
		log.Println("ALLOU YOBA ETO TI?")
		log.Println(msg)
		if err != nil {
			log.Fatal(constants.ErrorReadingMessage)
		}
	}
}
