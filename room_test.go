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

func TestEnter(t *testing.T) {
	room := Room{
		name:    "room",
		members: []Member{},
		history: &HistoryTest{},
	}

	m := &MemberTestImpl{
		name: "hoge",
	}

	err := room.Enter(m)
	if err != nil {
		t.Errorf("err :%v", err)
	}

	if len(room.members) != 1 {
		t.Errorf("not expected members num\n e:%d\n a:%d", 1, len(room.members))
	}
	a := room.members[0]
	if a.Name() != "hoge" {
		t.Errorf("not expected member\n e:%s\n a:%s", "hoge", a.Name())
	}
}

func TestExit(t *testing.T) {
	room := Room{
		name: "room",
		members: []Member{
			&MemberTestImpl{name: "hoge"},
			&MemberTestImpl{name: "fuga"},
		},
		history: &HistoryTest{},
	}

	err := room.Exit(&MemberTestImpl{name: "hoge"})
	if err != nil {
		t.Error(err)
	}

	if len(room.members) != 1 {
		t.Errorf("not expected members num\n e:%d\n a:%d", 1, len(room.members))
	}
	a := room.members[0]
	if a.Name() != "fuga" {
		t.Errorf("not expected member\n e:%s\n a:%s", "fuga", a.Name())
	}
}

func TestName(t *testing.T) {
	room := Room{
		name: "room",
	}

	if room.Name() != "room" {
		t.Errorf("not expected name\n e:%s\n a:%s\n", "room", room.Name())
	}
}
