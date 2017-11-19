package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type Hub struct {
	Rooms []Roomer
}

func (h *Hub) ExtractRoom(name string) (Roomer, error) {
	if len(h.Rooms) == 0 {
		return createNewRoom(h, name)
	}

	for _, r := range h.Rooms {
		if name == r.NameStr() {
			return r, nil
		}
	}

	return createNewRoom(h, name)
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

func createNewRoom(h *Hub, name string) (Roomer, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "failed create log file")
	}
	history, err := NewHistory(filepath.Join(pwd, "history"), name)
	if err != nil {
		return nil, errors.Wrap(err, "failed create log file")
	}
	room := NewRoom(name, history)
	h.Rooms = append(h.Rooms, room)
	log.Printf("Create New Room: %s\n", name)
	return room, nil
}
