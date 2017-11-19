package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":8080", "server port")
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var hub Hub = Hub{}

func handler(w http.ResponseWriter, r *http.Request) {
	roomParam := r.URL.Query().Get("room")
	name := r.URL.Query().Get("name")

	if roomParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("necessary 'room' parameter"))
		log.Println("necessary 'room' parameter")
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(hub)
	room := hub.ExtractRoom(roomParam)
	if name == "" {
		name = fmt.Sprintf("hoge%d", room.Count()+1)
	}
	log.Printf("room: %s", room.NameStr())
	log.Printf("name: %s", name)

	client := NewClient(name, conn, room)
	room.Enter(client)

	go client.Socket()
}

func run() error {
	http.HandleFunc("/wc", handler)
	return http.ListenAndServe(*addr, nil)
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}
