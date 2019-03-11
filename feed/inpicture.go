package feed

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type InPictureFeed struct {
	BaseFeedHandler
}

const InPictureFeedName = "inpicture"

func (ip InPictureFeed) Handle(id string, url string, updated int64) (*Feed, error) {
	resp, err := ip.Parse(url)
	if err != nil {
		log.Printf("Unable to read and parse %s: %v\n", url, err)
		return nil, err
	}
	feed, err := ip.ParseFeed(url, id, updated, resp)
	if err != nil {
		log.Printf("Unable to parse inPicture feed %s from HTML: %v\n", url, err)
		return nil, err
	}
	return feed, nil
}

func (ip *InPictureFeed) ParseFeed(url string, id string, updated int64, doc *goquery.Document) (*Feed, error) {
	title := strings.TrimSpace(doc.Find(".content__headline").Text())
	subheader := strings.TrimSpace(doc.Find(".content__standfirst p").Text())
	log.Printf("TITLE: %s\n", title)
	log.Printf("SUBHEADER: %s\n", subheader)

	items := doc.Find(".gallery__item")
	feedImages := make([]FeedImage, items.Length())

	items.Each(func(index int, selection *goquery.Selection) {
		caption := strings.TrimSpace(selection.Find(".gallery__caption").Text())
		srcNode := selection.Find(".gallery__img-container picture img")
		url, _ := srcNode.Attr("src")

		feedImages[index] = FeedImage{Caption: caption, Image: url}
	})
	return NewFeed(InPictureFeedName, id, title, subheader, url, feedImages, updated), nil
}