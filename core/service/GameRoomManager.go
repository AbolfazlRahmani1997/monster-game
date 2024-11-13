package service

import (
	"Monstern/core/domain"
	"fmt"
)

type GameRoomManager struct {
	Players            []domain.Player
	GamesStart         []domain.Game
	GameWaitingToStart map[string]*domain.Game
	Waiting            chan domain.Player
}

func NewGameRoomManager() *GameRoomManager {
	return &GameRoomManager{
		Waiting:            make(chan domain.Player, 10),
		GameWaitingToStart: make(map[string]*domain.Game),
		Players:            make([]domain.Player, 0),
		GamesStart:         make([]domain.Game, 0),
	}
}

func (receiver *GameRoomManager) Watching() {

}

func (receiver *GameRoomManager) Start() {
	for {
		select {
		case player := <-receiver.Waiting:
			lenGameWaiting := len(receiver.GameWaitingToStart)
			if lenGameWaiting > 0 {

				for _, game := range receiver.GameWaitingToStart {

					game.JoinPlayer(player)
					fmt.Println(len(game.Players))
					if len(game.Players) == 2 {
						go game.Run()
						game.Pour <- true
					}
					break
				}

			} else {
				cards := domain.NewCollection()
				newGame := domain.NewGame(cards)
				receiver.GameWaitingToStart[newGame.Id] = newGame
				newGame.Players[player.Id] = player
				player.Message <- domain.Message{MessageType: "Join In Game", Data: newGame.Id}
			}

		}
	}

}

func (receiver *GameRoomManager) InvitePlayer(playerId string, name string) {
	player := domain.Player{Id: playerId, Name: name}
	receiver.Waiting <- player
}
