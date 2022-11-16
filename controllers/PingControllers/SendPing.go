package PingControllers

import (
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
)

func HandleSendPing(context *gin.Context) {
	node, _ := snowflake.NewNode(1)
	context.JSON(200, gin.H{
		"message": node.Generate(),
	})
}
