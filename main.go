package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	injectedHeaderServer := router.Group("/", func(c *gin.Context) {
		c.Header("Server", "Gin")
		c.Header("X-GIN-VERSION", gin.Version)

		// For middleware
		//c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Not Found"})
		//c.Abort()

		c.Next()

	})

	routes(injectedHeaderServer)

	err := router.Run(":" + ServerPort)

	if err != nil {
		panic(err)
	}

}
