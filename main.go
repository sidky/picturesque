package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H { "message": "pong" })
	})
	r.POST("/feed/update", func(c *gin.Context) {
		b, _ := ioutil.ReadAll(c.Request.Body)
		log.Print(string(b))
		c.JSON(200, gin.H { "message": "pong" })
	})
	r.Run()
}


