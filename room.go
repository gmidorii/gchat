package main

type Roomer interface {
	NameStr() string
	Count() int
	Message([]byte) error
	Enter(Clienter) error
	Exit(Clienter) error
}

type Room struct {
	Name    string
	clients []Client
}

func NewRoom(name string) Roomer {
	return &Room{
		Name:    name,
		clients: []Client{},
	}
}

func (r *Room) NameStr() string {
	return r.Name
}

func (r *Room) Count() int {
	return len(r.clients)
}

func (r *Room) Message(b []byte) error {
	for _, c := range r.clients {
		err := c.Send(b)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Room) Enter(c Clienter) error {
	return nil
}

func (r *Room) Exit(c Clienter) error {
	return nil
}
