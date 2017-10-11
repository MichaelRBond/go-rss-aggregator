package types

// RSSFeed describes an RSS feed in saved in the database
type RSSFeed struct {
	ID    int16
	Title string
	URL   string
}
