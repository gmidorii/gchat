package main

import (
	"fmt"
	"log"
)

type Roomer interface {
	NameStr() string
	Count() int
	Message([]byte, Member) error
	Enter(*MemberImpl) error
	Exit(*MemberImpl) error
}

type Room struct {
	name    string
	members []*MemberImpl
	history Historier
}

func NewRoom(name string, history Historier) Roomer {
	return &Room{
		name:    name,
		members: []*MemberImpl{},
		history: history,
	}
}

func (r *Room) NameStr() string {
	return r.name
}

func (r *Room) Count() int {
	return len(r.members)
}

func (r *Room) Message(b []byte, sender Member) error {
	mes := fmt.Sprintf("%s: %s \n %s",
		r.NameStr(),
		sender.Name(),
		string(b),
	)

	for _, m := range r.members {
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

func (r *Room) Enter(m *MemberImpl) error {
	r.members = append(r.members, m)
	log.Printf("%s entered to %s", m.Name(), r.name)
	return nil
}

func (r *Room) Exit(m *MemberImpl) error {
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
