package PingControllers

import "github.com/gin-gonic/gin"

func HandleSendPing(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}
