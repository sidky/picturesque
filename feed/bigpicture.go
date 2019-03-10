package feed

import (
	"github.com/PuerkitoBio/goquery"
	"log"
)

type BigPictureFeed struct {
	BaseFeedHandler
}

func (b *BigPictureFeed) Handle(url string) (*Feed, error) {
	resp, err := b.Parse(url)
	if err != nil {
		log.Printf("Unable to read  and parse %s: %v\n", url, err)
		return nil, err
	}
	feed, err := b.ParseFeed(resp)
	if err != nil {
		log.Printf("Unable to parse %s from HTML: %v", url, err)
		return nil, err
	}
	return feed, nil
}

func (b *BigPictureFeed) ParseFeed(doc *goquery.Document) (*Feed, error) {

	container := doc.Find("body").Find("div #container")

	title := container.Find("h3").Text()
	subhead := container.Find(".subhead").Text()

	photos := container.Find("div .photo").Find("img")
	caption := container.Find(".pcaption").Find(".gcaption")

	urls := make([]string, photos.Length())
	photoCaptions := make([]string, caption.Length())

	photos.Each(func (index int, node *goquery.Selection) {
		urls[index], _ = node.Attr("src")
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

	return &Feed{
		Title: title,
		Description: subhead,
		Images: feedImages,
	}, nil
}

