package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes(server)

	err := server.Run(":" + ServerPort)

	if err != nil {
		panic(err)
	}

}
