package domain

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Player struct {
	Id      string
	Name    string
	conn    websocket.Conn
	Message chan string
}

func (receiver *Player) Read() {
	for {
		_, message, err := receiver.conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println(string(message))
	}
}
func (c *Player) writeMessage() {
	defer func() {

	}()

	for {
		message, ok := <-c.Message
		if !ok {
			err := c.conn.WriteJSON(message)
			if err != nil {
				return
			}

		}

	}
}
