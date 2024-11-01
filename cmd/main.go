package main

import (
	"Monstern/core/domain"
	"fmt"
)

func main() {
	c := domain.NewCollection()

	game := domain.NewGame(c)
	game.JoinPlayer(domain.Player{Id: "10", Name: "Test"})
	game.JoinPlayer(domain.Player{Id: "11", Name: "Ali"})
	game.JoinPlayer(domain.Player{Id: "12", Name: "Ho"})
	game.JoinPlayer(domain.Player{Id: "13", Name: "i"})

	go game.Run()
	game.Pour <- true
	for {
		var player string
		var card string
		_, err := fmt.Scanln(&player, &card)
		if err != nil {
			return
		}
		if card == "check" {
			game.Check <- player
		}
		game.Pick <- domain.PickCard{Player: player, Card: card}
	}
}
