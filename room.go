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
	Name    string
	members []*MemberImpl
}

func NewRoom(name string) Roomer {
	return &Room{
		Name:    name,
		members: []*MemberImpl{},
	}
}

func (r *Room) NameStr() string {
	return r.Name
}

func (r *Room) Count() int {
	return len(r.members)
}

func (r *Room) Message(b []byte, sender Member) error {
	m := fmt.Sprintf("%s: %s \n %s",
		r.NameStr(),
		sender.Name(),
		string(b),
	)

	for _, c := range r.members {
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

func (r *Room) Enter(m *MemberImpl) error {
	r.members = append(r.members, m)
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

	log.Printf("%s is exited from %s", m.Name(), r.NameStr())
	return nil
}
