package domain

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/oklog/ulid/v2"
)

const (
	InGame  Status = "InGame"
	Waiting Status = "Waiting"
)

type Status string

type Player struct {
	Id      string
	Name    string
	Conn    *websocket.Conn
	Message chan Message
	Pick    chan PickCard
	Game    *Game
	Status  Status
}

func NewPlayer(name string, Conn *websocket.Conn) *Player {
	return &Player{Id: ulid.Make().String(), Conn: Conn, Name: name, Message: make(chan Message, 10)}
}

func (receiver *Player) Read() {

	for {
		var Receive Message
		_, message, err := receiver.Conn.ReadMessage()
		err = json.Unmarshal(message, &Receive)
		if err != nil {
			return
		}
		if err != nil {
			return
		}
		fmt.Print(Receive)
		switch Receive.MessageType {
		case "pick":
			receiver.Pick <- PickCard{Player: receiver.Id, Card: Receive.CardId}
		case "Pour":
			receiver.Game.Pour <- true
		case "check":
			receiver.Game.Check <- receiver.Id

		default:
		}

	}
}
func (receiver *Player) WriteMessage() {
	defer func() {
		receiver.Conn.Close()
	}()

	for {
		message, ok := <-receiver.Message
		if ok {
			err := receiver.Conn.WriteJSON(message)
			if err != nil {
				return
			}

		}

	}
}
