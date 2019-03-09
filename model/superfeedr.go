package model

type FeedStatus struct {
	Status string `json:"feed"`
	Title string `json:"title"`
	HttpCode int32 `json:"code"`
	Http string `json:"http"`
	NextFetch int64 `json:"nextFetch"`
	Period int32 `json:"period"`
	LastFetch int64 `json:"lastFetch"`
	LastParse int64 `json:"lastParse"`
	LastMaintenanceAt int64 `json:"lastMaintenanceAt"`
	EntriesCount int32 `json:"entriesCountSinceLastMaintenance"`
	Velocity float64 `json:"velocity"`
	Popularity float64 `json:"popularity"`
}

type FeedLink struct {
	Href string `json:"href"`
	Rel *string `json:"rel"`
	MIMEType *string `json:"type"`
	Title *string `json:"title"`
}

type Category struct {
	Term string `json:"term"`
}

type FeedObject struct {
	ID *string `json:"id"`
	Title *string `json:"title"`
	Published *int64 `json:"published"`
	Updated *int64 `json:"updated"`
	Content *string `json:"content"`
	Permalink *string `json:"permalinkUrl"`
	Links map[string][]FeedLink `json:"standardLinks"`
}

type FeedEntry struct {
	FeedObject
	Summary *string `json:"summary"`
	Source *FeedObject `json:"source"`
	Language *string `json:"language"`
}

type FeedCallback struct {
	Status FeedStatus `json:"status"`
	Title *string `json:"title"`
	Items []FeedEntry `json:"items"`
}
