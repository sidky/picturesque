package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H { "message": "pong" })
	})
	r.POST("/feed/update", func(c *gin.Context) {
		log.Print(c.Request.Body)
		c.JSON(200, gin.H { "message": "pong" })
	})
	port := os.Getenv("PORT")
	r.Run()
}


