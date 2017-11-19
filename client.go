package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Clienter interface {
	Send([]byte, *Client) error
	Socket()
}

type Client struct {
	HandleName string
	Conn       *websocket.Conn
	Room       Roomer
}

func NewClient(name string, conn *websocket.Conn, room Roomer) *Client {
	return &Client{
		HandleName: name,
		Conn:       conn,
		Room:       room,
	}
}

func (c *Client) Send(b []byte, sender *Client) error {
	response := fmt.Sprintf("%s: %s \n %s",
		c.Room.NameStr(),
		sender.HandleName,
		string(b),
	)
	// mtype?
	err := c.Conn.WriteMessage(1, []byte(response))
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) Socket() {
	for {
		mtype, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("%d: %s", mtype, string(p))
		response := fmt.Sprintf("%s:%s \n %s",
			c.Room.NameStr(),
			c.HandleName,
			string(p),
		)
		c.Room.Message(p, c)

		if err := c.Conn.WriteMessage(mtype, []byte(response)); err != nil {
			log.Println(err)
			return
		}
	}
}
