package api

import (
	"Monstern/adapters/handler"
	"github.com/gin-gonic/gin"
)

var engin *gin.Engine

func Init(gameHandler handler.GameHandler) {
	engin = gin.Default()
	engin.GET("/", gameHandler.RegisterInGame)

}

func Start(addr string) error {
	return engin.Run(addr)
}
