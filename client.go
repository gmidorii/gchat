package main

type Clienter interface {
	Send([]byte) error
}

type Client struct {
	HandleName string
}

func NewClient(name string) Clienter {
	return &Client{
		HandleName: name,
	}
}

func (c *Client) Send(b []byte) error {
	return nil
}
