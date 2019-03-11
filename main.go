package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"picturesque/feed"
	"picturesque/model"
	"picturesque/storage"
)

func parseSuperfeedrCallback(req *http.Request) (*model.FeedCallback, error) {
	defer func() {
		if err := req.Body.Close(); err != nil {
			log.Print("Error closing request body: %v\n", err)
		}
	}()
	b, _ := ioutil.ReadAll(req.Body)
	cb := model.FeedCallback{}
	e := json.Unmarshal(b, &cb)
	if e != nil {
		return nil, e
	}
	return &cb, nil
}

func handleSuperfeedrCallback(c *gin.Context, handler feed.FeedHandler, feedStore *storage.FirebaseStore) {
	cb, err := parseSuperfeedrCallback(c.Request)
	if err != nil {
		log.Print("Unable to parse callback: %v\n", err)
		c.JSON(400, gin.H{"message": "Unable to parse"})
		return
	}

	for _, item := range cb.Items {
		id := *item.ID
		updated := *item.Updated
		feed, err := handler.Handle(id,*item.Permalink, updated)

		if err != nil {
			log.Printf("Unable to load item %s: %v\n", *item.ID, err)
			c.JSON(500, gin.H{"message": "Unable to load"})
			return
		}

		err = feedStore.AddFeed(feed)
		if err != nil {
			log.Printf("Unable to store in firebase: %v\n", err)
			c.JSON(500, gin.H{"message": "Firebase error"})
			return
		}
	}
	c.JSON(200, gin.H{"message": "successful"})
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H { "message": "pong" })
	})

	bigPicture := feed.BigPictureFeed{}
	inFocus := feed.InFocusFeed{}
	feedStore := storage.FirebaseStore{}
	feedStore.Init()

	r.POST("/feed/update/bigpicture", func(c *gin.Context) {
		handleSuperfeedrCallback(c, bigPicture, &feedStore)
	})

	r.POST("/feed/update/infocus", func (c *gin.Context) {
		handleSuperfeedrCallback(c, inFocus, &feedStore)
	})
	r.POST("/feed/update/inpicture", func (c *gin.Context) {
		b, _ := ioutil.ReadAll(c.Request.Body)
		cb := model.FeedCallback{}
		json.Unmarshal(b, &cb)

		r, _ := json.Marshal(cb)
		log.Print(string(r))
	})

	r.Run()
}


