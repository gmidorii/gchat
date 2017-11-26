package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

type Roomer interface {
	Name() string
	Count() int
	Message([]byte, Member) error
	Enter(Member) error
	Exit(Member) error
}

type Room struct {
	name    string
	members []Member
	history Historier
}

func NewRoom(name string, root string) (Roomer, error) {
	history, err := NewHistory(root, name)
	if err != nil {
		return nil, errors.Wrap(err, "failed create log file")
	}
	return &Room{
		name:    name,
		members: []Member{},
		history: history,
	}, nil
}

func (r *Room) Name() string {
	return r.name
}

func (r *Room) Count() int {
	return len(r.members)
}

func (r *Room) Message(b []byte, sender Member) error {
	mes := fmt.Sprintf("%s: %s \n %s",
		r.Name(),
		sender.Name(),
		string(b),
	)

	for _, m := range r.members {
		log.Println(m)
		if sender.Name() == m.Name() {
			continue
		}
		err := m.Send(mes)
		if err != nil {
			return err
		}
	}
	// write log
	return r.history.Write(r.name, sender.Name(), string(b))
}

func (r *Room) Enter(m Member) error {
	r.members = append(r.members, m)
	log.Printf("%s entered to %s", m.Name(), r.name)
	return nil
}

func (r *Room) Exit(m Member) error {
	var index int
	for i, member := range r.members {
		if m.Name() == member.Name() {
			index = i
			break
		}
	}
	r.members = append(r.members[:index], r.members[index+1:]...)
	log.Printf("%s exited from %s", m.Name(), r.name)

	if len(r.members) == 0 {
		return r.close()
	}
	return nil
}

func (r *Room) close() error {
	err := r.history.Close()
	if err != nil {
		return err
	}
	hub.Close(r)
	return nil
}
