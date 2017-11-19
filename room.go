package main

import "fmt"

type Roomer interface {
	NameStr() string
	Count() int
	Message([]byte, Clienter) error
	Enter(*Client) error
	Exit(*Client) error
}

type Room struct {
	Name    string
	clients []*Client
}

func NewRoom(name string) Roomer {
	return &Room{
		Name:    name,
		clients: []*Client{},
	}
}

func (r *Room) NameStr() string {
	return r.Name
}

func (r *Room) Count() int {
	return len(r.clients)
}

func (r *Room) Message(b []byte, sender Clienter) error {
	m := fmt.Sprintf("%s: %s \n %s",
		r.NameStr(),
		sender.Name(),
		string(b),
	)

	for _, c := range r.clients {
		if sender.Name() == c.HandleName {
			continue
		}
		err := c.Send(m)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Room) Enter(c *Client) error {
	r.clients = append(r.clients, c)
	return nil
}

func (r *Room) Exit(c *Client) error {
	return nil
}
