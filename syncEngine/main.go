package syncEngine

import (
	"github.com/michaelrbond/go-rss-aggregator/logger"
	"github.com/michaelrbond/go-rss-aggregator/types"
	"github.com/michaelrbond/go-rss-aggregator/utils"
)

// SyncRssFeeds initiates the download
func SyncRssFeeds(context *types.Context) {
	logger.Debug("Executing RSS Sync.")

	feeds, err := utils.RSSGetAllFeedsFromDB(context)
	if err != nil {
		return
	}

	for _, feed := range feeds {

		feedContent, err := feed.Download()
		if err != nil {
			continue
		}

		feed.Process(context, feedContent)
	}
}
