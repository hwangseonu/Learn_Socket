package main

import (
	"github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func main() {
	port := ":5000"
	server, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}
	server.On("connection", func (so socketio.Socket) {
		log.Printf("connected from %s", so.Request().RemoteAddr)
		so.On("chat message", func(msg string) {
			so.Emit("chat message", msg)
			log.Printf("chat: %s", msg)
			so.BroadcastTo("chat", "chat message", msg)
		})
		so.On("disconnection", func () {
			log.Printf("disconnected from %s", so.Request().RemoteAddr)
		})
	})
	server.On("error", func (so socketio.Socket, err error) {
		log.Fatal(err)
	})
	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}