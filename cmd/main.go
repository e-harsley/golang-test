package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-play/api/router"
	"github.com/go-play/config"
	"github.com/go-play/database"
)

func main() {
	server := gin.Default()

	database.Init()

	router.Init(server)

	port := config.PORT

	server.Run(":" + port)
}
