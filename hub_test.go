package main

import "testing"

func TestExtractRoom(t *testing.T) {
	cases := []struct {
		rooms []Roomer
		in    string
		want  Roomer
	}{
		{[]Roomer{}, "hoge", &Room{name: "hoge"}},
	}

	for _, c := range cases {
		hub := Hub{
			Rooms: c.rooms,
		}
		r, err := hub.ExtractRoom(c.in)
		if err != nil {
			t.Errorf("failed extract: %v", err)
		}
		if c.want.NameStr() != r.NameStr() {
			t.Errorf("not expected room\n e:%s \n a:%s", c.want.NameStr(), r.NameStr())
		}
	}
}
