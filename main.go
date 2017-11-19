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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	roomParam := r.URL.Query().Get("room")
	name := r.URL.Query().Get("name")

	if roomParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("necessary 'room' parameter"))
		return
	}
	room := hub.ExtractRoom(roomParam)
	if name == "" {
		name = fmt.Sprintf("hoge%d", room.Count()+1)
	}

	client := NewClient(name)
	room.Enter(client)

	for {
		mtype, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%d: %s", mtype, string(p))
		if err := conn.WriteMessage(mtype, p); err != nil {
			log.Println(err)
			return
		}
	}
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
