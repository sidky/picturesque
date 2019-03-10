package storage

import "google.golang.org/api/option"

type FirebaseStore struct { }

func (f *FirebaseStore) init() {
	option.WithCredentialsJSON()
}
