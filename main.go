package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"picturesque/feed"
	"picturesque/model"
	"picturesque/storage"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H { "message": "pong" })
	})

	bigPicture := feed.BigPictureFeed{}
	feedStore := storage.FirebaseStore{}
	feedStore.Init()

	r.POST("/feed/update/bigpicture", func(c *gin.Context) {
		b, _ := ioutil.ReadAll(c.Request.Body)
		log.Print(string(b))
		cb := model.FeedCallback{}
		e := json.Unmarshal(b, &cb)
		if e != nil {
			log.Print(e)
			c.JSON(400, gin.H { "message": "pong" })
		} else {
			for _, item := range cb.Items {
				id := *item.ID
				updated := *item.Updated
				feed, err := bigPicture.Handle(id, *item.Permalink, updated)
				if err != nil {
					log.Print("Unable to load: %v\n", err)
					c.JSON(500, gin.H{"message": "Unable to load"})
				} else {
					err = feedStore.AddFeed(feed)
					if err != nil {
						log.Print("Unable to store: %v\n", feed)
						c.JSON(500, gin.H{"message": "Firebase error"})
					}
				}
			}
			mb, _ := json.Marshal(cb)
			log.Print(string(mb))
			c.JSON(200, gin.H { "message": "pong" })
		}
	})

	r.POST("/feed/update/infocus", func (c *gin.Context) {
		b, _ := ioutil.ReadAll(c.Request.Body)
		log.Print(string(b))

		c.JSON(200, gin.H{"message": "successful"})
	})
	r.Run()
}


