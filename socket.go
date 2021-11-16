package main

import (
	"log"

	socketio "github.com/googollee/go-socket.io"
)

func Socket() *socketio.Server {

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")

		// so.Join("message")

		so.On("message", func(msg string) {
			so.Emit("message", "Return :- "+msg)
			log.Println("emit:", msg)
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	return server
}
