package feed

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"strings"
)

type BigPictureFeed struct {
	BaseFeedHandler
}

const FeedName = "bigpicture"

func (b BigPictureFeed) Handle(id string, url string, updated int64) (*Feed, error) {
	resp, err := b.Parse(url)
	if err != nil {
		log.Printf("Unable to read  and parse %s: %v\n", url, err)
		return nil, err
	}
	feed, err := b.ParseFeed(url, id, updated, resp)
	if err != nil {
		log.Printf("Unable to parse %s from HTML: %v", url, err)
		return nil, err
	}
	return feed, nil
}

func (b *BigPictureFeed) ParseFeed(url string, id string, updated int64, doc *goquery.Document) (*Feed, error) {

	container := doc.Find("body").Find("div #container")

	title := container.Find("h3").Text()
	subhead := container.Find(".subhead").Text()

	photos := container.Find("div .photo").Find("img")
	caption := container.Find(".pcaption").Find(".gcaption")

	urls := make([]string, photos.Length())
	photoCaptions := make([]string, caption.Length())

	photos.Each(func (index int, node *goquery.Selection) {
		url, _ := node.Attr("src")
		if strings.HasPrefix(url, "//") {
			url = "https:" + url
		}
		urls[index] = url
	})

	caption.Each(func (index int, node *goquery.Selection) {
		photoCaptions[index] = node.Text()
	})

	feedImages := make([]FeedImage, len(urls))

	for index, url := range urls {
		var caption string
		if index < len(photoCaptions) {
			caption = photoCaptions[index]
		} else {
			caption = ""
		}
		img := FeedImage{Image: url, Caption:caption}
		feedImages[index] = img
	}

	return NewFeed(FeedName, id, title, subhead, url, feedImages, updated), nil
}

