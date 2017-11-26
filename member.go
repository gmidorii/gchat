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
	handleName string
	conn       *websocket.Conn
	room       Roomer
}

func NewMember(name string, conn *websocket.Conn, room Roomer) Member {
	return &MemberImpl{
		handleName: name,
		conn:       conn,
		room:       room,
	}
}

func (m *MemberImpl) Name() string {
	return m.handleName
}

func (m *MemberImpl) Send(mes string) error {
	// mtype?
	err := m.conn.WriteMessage(1, []byte(mes))
	if err != nil {
		return err
	}
	return nil
}

func (m *MemberImpl) Socket() {
	defer func() {
		m.room.Exit(m)
	}()

	for {
		mtype, p, err := m.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		response := fmt.Sprintf("%s:%s \n %s",
			m.room.NameStr(),
			m.handleName,
			string(p),
		)
		m.room.Message(p, m)

		if err := m.conn.WriteMessage(mtype, []byte(response)); err != nil {
			log.Println(err)
			return
		}
	}
}
