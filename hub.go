package main

import (
	"log"
)

type Hub struct {
	Rooms       []Roomer
	historyRoot string
}

func NewHub(root string) *Hub {
	return &Hub{
		historyRoot: root,
	}
}

func (h *Hub) ExtractRoom(name string) (Roomer, error) {
	if len(h.Rooms) == 0 {
		room, err := NewRoom(name, h.historyRoot)
		if err != nil {
			return nil, err
		}
		h.Rooms = append(h.Rooms, room)
		log.Printf("Create New Room: %s\n", name)

		return room, nil
	}

	for _, r := range h.Rooms {
		if name == r.NameStr() {
			return r, nil
		}
	}

	room, err := NewRoom(name, h.historyRoot)
	if err != nil {
		return nil, err
	}
	h.Rooms = append(h.Rooms, room)
	log.Printf("Create New Room: %s\n", name)

	return room, nil
}

func (h *Hub) Close(r Roomer) {
	var index int
	for i, room := range h.Rooms {
		if r.NameStr() == room.NameStr() {
			index = i
			break
		}
	}
	h.Rooms = append(h.Rooms[:index], h.Rooms[index+1:]...)
	log.Printf("[%s] room is closed", r.NameStr())
}
