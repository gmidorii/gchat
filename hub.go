package main

type Hub struct {
	Rooms []Roomer
}

func (h *Hub) ExtractRoom(name string) Roomer {
	if len(h.Rooms) == 0 {
		room := NewRoom(name)
		h.Rooms = append(h.Rooms, room)
		return room
	}

	for _, r := range h.Rooms {
		if name == r.NameStr() {
			return r
		}
	}

	room := NewRoom(name)
	h.Rooms = append(h.Rooms, room)
	return room
}
