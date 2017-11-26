package main

import (
	"os"
	"testing"
)

type MemberTestImpl struct {
	name    string
	content string
}

func (m *MemberTestImpl) Name() string {
	return m.name
}

func (m *MemberTestImpl) Send(mes string) error {
	m.content = mes
	return nil
}

func (m *MemberTestImpl) Socket() {
}

func TestMessage_Content(t *testing.T) {
	exp := "room: sender \n content"
	room, _ := NewRoom("room", "./")
	sender := &MemberTestImpl{
		name: "sender",
	}
	room.Enter(sender)
	reciver := &MemberTestImpl{
		name: "reciver",
	}
	room.Enter(reciver)

	if err := room.Message([]byte("content"), sender); err != nil {
		t.Errorf("failed Message func: %v", err)
	}

	if reciver.content != exp {
		t.Errorf("not expected content\n e:%s\n a:%s\n", exp, reciver.content)
	}

	os.Remove("room")
}
