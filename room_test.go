package main

import (
	"fmt"
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

type HistoryTest struct {
	content string
}

func (h *HistoryTest) Write(room, member, message string) error {
	h.content = fmt.Sprintf("%s:%s:%s", room, member, message)
	return nil
}

func (h *HistoryTest) Close() error {
	return nil
}

func TestMessage_Content(t *testing.T) {
	// setup
	exp := "room: sender \n content"
	room := Room{
		name:    "room",
		members: []Member{},
		history: &HistoryTest{},
	}
	sender := &MemberTestImpl{
		name: "sender",
	}
	room.Enter(sender)
	reciver := &MemberTestImpl{
		name: "reciver",
	}
	room.Enter(reciver)

	// execute
	err := room.Message([]byte("content"), sender)
	if err != nil {
		t.Errorf("failed Message func: %v", err)
	}

	// verify
	if reciver.content != exp {
		t.Errorf("not expected content\n e:%s\n a:%s\n", exp, reciver.content)
	}
	h, _ := room.history.(*HistoryTest)
	if h.content != "room:sender:content" {
		t.Errorf("not expected content\n a:%s\n", h.content)
	}

	// tear down
	os.Remove("room")
}
