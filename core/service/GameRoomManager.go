package service

import "Monstern/core/domain"

type GameRoomManager struct {
	Players            []domain.Player
	GamesStart         []domain.Game
	GameWaitingToStart map[string]*domain.Game
	Waiting            chan domain.Player
}

func NewGameRoomManager() *GameRoomManager {
	return &GameRoomManager{}
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
					game.Players[player.Id] = player
					break
				}

			} else {
				cards := domain.NewCollection()
				newGame := domain.NewGame(cards)
				receiver.GameWaitingToStart[newGame.Id] = newGame
				newGame.Players[player.Id] = player

			}

		}
	}

}

func (receiver *GameRoomManager) InvitePlayer(playerId string, name string) {
	player := domain.Player{Id: playerId, Name: name}
	receiver.Waiting <- player

}
