package feed

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type FeedImage struct {
	Image string
	Caption string
}

func (f FeedImage) String() string {
	return fmt.Sprintf("{Image: %s, Caption: %s}", f.Image, f.Caption)
}

type Feed struct {
	FeedName string
	ID string
	Title string
	SubHeader string
	URL string
	Images []FeedImage
	UpdatedAt int64
}

func (f Feed) String() string {
	return fmt.Sprintf("{FeedName: %s, ID: %s, Title: %s, SubHeader: %s, URL: %s, Images: %v, UpdatedAt: %d}",
		f.FeedName, f.ID, f.Title, f.SubHeader, f.URL, f.Images, f.UpdatedAt)
}

func NewFeed(feedName string, id string, title string, subHeader string, url string, images []FeedImage, updatedAt int64) *Feed {
	return &Feed{
		FeedName: feedName,
		ID: id,
		Title: title,
		SubHeader: subHeader,
		URL: url,
		Images: images,
		UpdatedAt: updatedAt,
	}
}

type FeedHandler interface {
	Handle(id string, url string, updated int64) (*Feed, error)
}

type BaseFeedHandler struct { }

func (f *BaseFeedHandler) Parse(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Printf("Unable to close response body: %v\n", err)
		}
	}()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}
