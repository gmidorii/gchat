package main

import "log"

type Hub struct {
	Rooms []Roomer
}

func (h *Hub) ExtractRoom(name string) Roomer {
	if len(h.Rooms) == 0 {
		return createNewRoom(h, name)
	}

	for _, r := range h.Rooms {
		if name == r.NameStr() {
			return r
		}
	}

	return createNewRoom(h, name)
}

func createNewRoom(h *Hub, name string) Roomer {
	room := NewRoom(name)
	h.Rooms = append(h.Rooms, room)
	log.Printf("Create New Room: %s\n", name)
	return room
}
