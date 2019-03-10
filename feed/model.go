package feed

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type FeedImage struct {
	Image string
	Caption string
}

type Feed struct {
	Title string
	Description string
	URL string
	Images []FeedImage
}

type FeedHandler interface {
	handle(url string)
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
