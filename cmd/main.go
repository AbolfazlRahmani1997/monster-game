package main

import (
	"Monstern/adapters/handler"
	"Monstern/api"
	"Monstern/core/service"
)

func main() {
	gameRoomManager := service.NewGameRoomManager()
	go gameRoomManager.Start()
	gameHandler := handler.NewGameHandler(gameRoomManager)
	api.Init(*gameHandler)
	err := api.Start("0.0.0.0:9090")
	if err != nil {
		return
	}
}
