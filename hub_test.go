package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractRoom(t *testing.T) {
	cases := []struct {
		rooms []Roomer
		in    string
		want  Roomer
	}{
		{[]Roomer{}, "hoge", &Room{name: "hoge"}},
		{[]Roomer{&Room{name: "hoge"}}, "hoge", &Room{name: "hoge"}},
		{[]Roomer{&Room{name: "hoge"}}, "new", &Room{name: "new"}},
	}

	// set up
	root, err := os.Getwd()
	if err != nil {
		t.Errorf("failed pwd: %v", err)
	}
	for _, c := range cases {
		hub := Hub{
			Rooms:       c.rooms,
			historyRoot: root,
		}
		// execute
		r, err := hub.ExtractRoom(c.in)
		if err != nil {
			t.Errorf("failed extract: %v", err)
		}
		if c.want.NameStr() != r.NameStr() {
			t.Errorf("not expected room\n e:%s \n a:%s", c.want.NameStr(), r.NameStr())
		}

		// tear down
		_, err = os.Stat(filepath.Join(hub.historyRoot, r.NameStr()))
		if os.IsNotExist(err) {
			continue
		}

		err = os.Remove(filepath.Join(hub.historyRoot, r.NameStr()))
		if err != nil {
			t.Errorf("failed rm history file: %v", err)
		}
	}
}
