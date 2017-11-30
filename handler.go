package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var hub *Hub

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

	room, err := hub.ExtractRoom(roomParam)
	if err != nil {
		log.Println(err)
		return
	}
	if name == "" {
		name = fmt.Sprintf("gchat-%d", room.Count()+1)
	}

	member := NewMember(name, conn, room)
	room.Enter(member)

	go member.Socket()
}
