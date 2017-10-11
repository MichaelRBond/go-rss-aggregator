package syncEngine

import (
	"fmt"
	"net/http"

	"github.com/mmcdole/gofeed"

	"github.com/michaelrbond/go-rss-aggregator/logger"
	"github.com/michaelrbond/go-rss-aggregator/types"
	"github.com/michaelrbond/go-rss-aggregator/utils"
)

// SyncRssFeeds initiates the download
func SyncRssFeeds(context *types.Context) {
	logger.Debug("Executing RSS Sync.")

	// get all of the feeds in the database
	feeds, err := utils.RSSGetAllFeedsFromDB(context)
	if err != nil {
		return
	}

	// for each feed
	for _, feed := range feeds {
		fmt.Printf("TESTING: %q\n", feed)
		feedContent, err := downloadRSSFeed(feed.URL)
		if err != nil {
			continue
		}
		fmt.Printf("%+v\n", feedContent)
		// -- save items to the database
		// -- update the last sync in the database table
	}
}

func downloadRSSFeed(url string) (*gofeed.Feed, error) {
	resp, err := http.Get(url)

	if err != nil {
		logger.Error(fmt.Sprintf("Downloading feed: %s", err.Error()))
		return nil, err
	}

	fp := gofeed.NewParser()
	feed, err := fp.Parse(resp.Body)

	if err != nil {
		logger.Error(fmt.Sprintf("Parsing Feed: %s", err.Error()))
		return nil, err
	}

	return feed, nil
}
