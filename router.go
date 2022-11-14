package main

import (
	"github.com/gin-gonic/gin"
	"mvcgo/controllers/PingControllers"
	"mvcgo/controllers/TodoControllers"
)

func routes(route *gin.Engine) {
	route.GET("/ping", PingControllers.HandleSendPing)

	// CRUD for TO-DOS
	route.GET("/todos", TodoControllers.HandleListTodos)
	route.POST("/todos", TodoControllers.HandleCreateTodo)
	route.GET("/todos/:id", TodoControllers.HandleGetTodoById)
}
