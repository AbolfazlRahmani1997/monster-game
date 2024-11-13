package handler

import (
	"Monstern/core/domain"
	"Monstern/core/service"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type GameHandler struct {
	GameRoomManager *service.GameRoomManager
}

func NewGameHandler(manager *service.GameRoomManager) *GameHandler {
	return &GameHandler{GameRoomManager: manager}
}

func (receiver GameHandler) RegisterInGame(c *gin.Context) {
	conn, _ := upgrader.Upgrade(c.Writer, c.Request, nil)
	player := domain.NewPlayer(c.Query("name"), conn)
	receiver.GameRoomManager.Waiting <- *player
	go player.WriteMessage()

	player.Read()
}

type RegisterInGameRequest struct {
	Name string `json:"name"`
}
