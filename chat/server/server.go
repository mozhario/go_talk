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
	listener, err := net.Listen("tcp", fmt.Sprint(":%s", server.Port))
	if err != nil {
		log.Fatal(constants.ErrorListenerFailedToStart)
	}

	conn, _ := listener.Accept()
	conn.Write([]byte(constants.ChatStarted))

	for {
		buff := make([]byte, config.ServerBufferSize)
		msg, err := conn.Read(buff[0:])
		fmt.Println("ALLOU YOBA ETO TI?")
		fmt.Println(buff)
		fmt.Println(msg)
		if err != nil {
			log.Fatal(constants.ErrorReadingMessage)
		}
		// conn.Write(msg)
	}
}
