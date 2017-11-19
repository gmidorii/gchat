package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Member interface {
	Name() string
	Send(m string) error
	Socket()
}

type MemberImpl struct {
	HandleName string
	Conn       *websocket.Conn
	Room       Roomer
}

func NewMember(name string, conn *websocket.Conn, room Roomer) *MemberImpl {
	return &MemberImpl{
		HandleName: name,
		Conn:       conn,
		Room:       room,
	}
}

func (m *MemberImpl) Name() string {
	return m.HandleName
}

func (m *MemberImpl) Send(mes string) error {
	// mtype?
	err := m.Conn.WriteMessage(1, []byte(mes))
	if err != nil {
		return err
	}
	return nil
}

func (m *MemberImpl) Socket() {
	defer func() {
		m.Room.Exit(m)
	}()

	for {
		mtype, p, err := m.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%d: %s", mtype, string(p))
		response := fmt.Sprintf("%s:%s \n %s",
			m.Room.NameStr(),
			m.HandleName,
			string(p),
		)
		m.Room.Message(p, m)

		if err := m.Conn.WriteMessage(mtype, []byte(response)); err != nil {
			log.Println(err)
			return
		}
	}
}
