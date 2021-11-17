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
		log.Println("New connection")

		so.Join("chat")

		so.On("message", func(msg string) {
			// return that message
			so.Emit("message", msg)

			so.BroadcastTo("chat", "message", msg)
		})

		so.On("disconnection", func() {
			log.Println("Connection closed")
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	return server
}
