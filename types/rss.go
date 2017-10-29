package types

import (
	"fmt"
	"net/http"
	"time"

	"github.com/michaelrbond/go-rss-aggregator/logger"
	"github.com/mmcdole/gofeed"
)

// RSSFeedBase describes an RSS Feed
type RSSFeedBase struct {
	Title string
	URL   string
}

// Save saves a feed to the database
func (feed RSSFeedBase) Save(context *Context) error {
	// TODO : check for duplicates - Need to set field to unique in the database
	_, err := context.Db.Query("INSERT INTO `feeds` (`title`, `url`) VALUES(?, ?);", feed.Title, feed.URL)
	if err != nil {
		logger.Error(fmt.Sprintf("Error saving new feed: %s", err.Error()))
		return err
	}

	return nil
}

// RSSFeed describes an RSS feed in saved in the database
type RSSFeed struct {
	RSSFeedBase
	ID          int16
	LastUpdated int32
}

// Save saves an RSSFeed to the database
func (feed RSSFeed) Save(context *Context) error {
	// Update the Feed
	return nil
}

// Download Downloads the RSS feed and returns its content
func (feed RSSFeed) Download() (*gofeed.Feed, error) {
	resp, err := http.Get(feed.URL)

	if err != nil {
		logger.Error(fmt.Sprintf("Downloading feed: %s", err.Error()))
		return nil, err
	}

	fp := gofeed.NewParser()
	feedContent, err := fp.Parse(resp.Body)

	if err != nil {
		logger.Error(fmt.Sprintf("Parsing Feed: %s", err.Error()))
		return nil, err
	}

	return feedContent, nil
}

// Process processes feed items
func (feed RSSFeed) Process(context *Context, feedContent *gofeed.Feed) error {

	fmt.Printf("Last Processed feed: %d\n", feed.LastUpdated)
	fmt.Printf("%s\n", feedContent.Title)

	for _, item := range feedContent.Items {
		fmt.Printf("\t%s\n", item.Title)
		fmt.Printf("\t\t%s -- %s\n", item.Published, item.Updated)
		fmt.Printf("\t\t%v -- %v\n", item.PublishedParsed, item.UpdatedParsed)
		// fmt.Printf("\t\tdescription: %s\n", item.Description)
		fmt.Printf("\t\tcontent: %s\n", item.Content)
		fmt.Printf("\t\tlink: %s\n", item.Link)
		fmt.Printf("\t\tAuthor: %s\n", item.Author)
		fmt.Printf("\t\tGUID: %s\n", item.GUID)
		// fmt.Printf("\t\t%s\n", );
		// fmt.Printf("\t\t%s\n", );
		// fmt.Printf("\t\t%s\n", );
		// fmt.Printf("\t\t%+v\n", item.Extensions)

		// -- determine object time
		itemTime := rssItemTime(item)
		fmt.Printf("Itemtime: %d\n", itemTime)
		if itemTime >= feed.LastUpdated {
			rssItem := feed.BuildRssItemBase(item)
			rssItem.Save(context)
		}
	}

	feed.LastUpdated = int32(time.Now().Unix())
	if err := feed.Save(context); err != nil {
		return err
	}
	return nil
}

// BuildRssItemBase returns a RSSItemBase from feed item
func (feed RSSFeed) BuildRssItemBase(item *gofeed.Item) RSSItemBase {
	base := RSSItemBase{}
	base.FeedID = feed.ID
	base.Title = item.Title
	base.Description = item.Description
	base.Content = item.Content
	base.Link = item.Link
	base.Updated = int32(item.UpdatedParsed.Unix())
	base.Published = int32(item.PublishedParsed.Unix())
	base.Author = item.Author.Name
	base.GUID = item.GUID
	base.Image = item.Image.URL
	// base.Categories = item.Categories // TODO :
	// base.Enclosures = item.Enclosure // TODO :
	return base
}

// RSSItemBase describes an item in a feed before it is saved to the database
type RSSItemBase struct {
	FeedID      int16
	Title       string
	Description string
	Content     string
	Link        string
	Updated     int32
	Published   int32
	Author      string
	GUID        string
	Image       string
	Categories  []string // TODO : Handle these properly
	Enclosures  string   // TODO : Handle these properly
}

// Save saves an RSSItemBase to the database
func (item RSSItemBase) Save(contect *Context) error {
	return nil
}

// RSSItem descriibes an item in a feed, saved in the database
type RSSItem struct {
	RSSItemBase
	ID      int32
	Read    bool
	Starred bool
}

// Save saves an RSS Item to the database
func (item RSSItem) Save(context *Context) error {

	// determine if this is new or an update
	// if new insert
	// if update, update

	return nil
}
