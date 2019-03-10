package feed

import (
	"github.com/PuerkitoBio/goquery"
	"log"
)

type InFocusFeed struct {
	BaseFeedHandler
}

const InFocusFeedName = "infocus"

func (inFocus InFocusFeed) Handle(id string, url string, updated int64) (*Feed, error) {
	resp, err := inFocus.Parse(url)
	if err != nil {
		log.Printf("Unable to read and parse infocus feed %s: %v\n", url, err)
		return nil, err
	}
	feed, err := inFocus.ParseFeed(url, id, updated, resp)
	if err != nil {
		log.Printf("Unable to parse infocus feed %s from HTML: %v\n", url, err)
		return nil, err
	}
	return feed, nil
}

func (inFocus *InFocusFeed) ParseFeed(url string, id string, updated int64, doc *goquery.Document) (*Feed, error) {
	title := doc.Find("#article-header h1").Text()
	subheader := doc.Find("#article-content p").Text()

	log.Printf("TITLE: %s\n", title)
	log.Printf("SUBHEADER: %s\n", subheader)

	photos := doc.Find(".photos .photo")
	log.Println(photos.Length())

	images := make([]FeedImage, photos.Length())

	photos.Each(func(index int, selection *goquery.Selection) {
		url, _ := selection.Find("picture source").First().Attr("data-srcset")
		caption := selection.Find(".caption span").Text()

		img := FeedImage{Image: url, Caption: caption}
		images[index] = img
	})

	return NewFeed(
		InFocusFeedName,
		id,
		title,
		subheader,
		url,
		images,
		updated), nil

}
