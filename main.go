package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"picturesque/model"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H { "message": "pong" })
	})
	r.POST("/feed/update/bigpicture", func(c *gin.Context) {
		a := c.GetHeader("Authorization")
		log.Printf("Auth: %s", a)
		b, _ := ioutil.ReadAll(c.Request.Body)
		log.Print(string(b))
		cb := model.FeedCallback{}
		e := json.Unmarshal(b, &cb)
		if e != nil {
			log.Print(e)
			c.JSON(400, gin.H { "message": "pong" })
		} else {
			mb, _ := json.Marshal(cb)
			log.Print(string(mb))
			c.JSON(200, gin.H { "message": "pong" })
		}
	})
	r.Run()
}


