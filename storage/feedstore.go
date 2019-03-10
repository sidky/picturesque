package storage

import (
	"context"
	"firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
	"log"
	"os"
	"picturesque/feed"
	"strconv"
	"strings"
)

type FirebaseStore struct {
	app *firebase.App
	dbClient *db.Client
}

func (f *FirebaseStore) Init() {
	b := []byte(os.Getenv("FIREBASE_ACCOUNT"))
	opt := option.WithCredentialsJSON(b)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing firebase: %v\n", err)
	}

	f.app = app

	client, err := f.app.DatabaseWithURL(context.Background(), os.Getenv("FIREBASE_HOST"))
	if err != nil {
		log.Fatalf("Error creating database client: %v\n", err)
	}
	f.dbClient = client
}

func (f *FirebaseStore) AddFeed(feed *feed.Feed) error {
	ref := f.dbClient.NewRef(feed.FeedName)
	entryRoot := ref.Child(validId(feed.ID))

	ctx := context.Background()

	err := entryRoot.Update(ctx, map[string]interface{} {
		"title": feed.Title,
		"subheader": feed.SubHeader,
		"url": feed.URL,
		"updated": feed.UpdatedAt,
	})
	if err != nil {
		log.Fatalf("Unable to update feed: %v\n", err)
		return err
	}

	imageRoot := entryRoot.Child("images")

	for n, img := range feed.Images {
		indexRoot := imageRoot.Child(strconv.Itoa(n))

		err := indexRoot.Update(ctx, map[string]interface{} {
			"image": img.Image,
			"caption": img.Caption,
		})

		if err != nil {
			log.Fatalf("Unable to update image: %v\n", err)
			return err
		}
	}
	return nil
}

func validId(id string) string {
	invalids := []string{"\\", "/", ".", "$", "#", "[", "]"}
	for _, r := range invalids {
		id = strings.ReplaceAll(id, r, "_")
	}
	return id
}

