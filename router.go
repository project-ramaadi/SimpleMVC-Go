package main

import (
	"github.com/gin-gonic/gin"
	"mvcgo/controllers/PingControllers"
	"mvcgo/controllers/TodoControllers"
)

func routes(route *gin.Engine) {
	route.GET("/ping", PingControllers.HandleSendPing)
	route.GET("/todos", TodoControllers.HandleListTodos)
}
