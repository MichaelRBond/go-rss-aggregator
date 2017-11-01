package utils

import (
	"fmt"

	"github.com/michaelrbond/go-rss-aggregator/logger"
	"github.com/michaelrbond/go-rss-aggregator/types"
)

// RSSGetAllFeedsFromDB returns an array of all the saved RSS Feeds
func RSSGetAllFeedsFromDB(context *types.Context) ([]types.RSSFeed, error) {
	result, err := context.Db.Query("SELECT * FROM `feeds`;")
	if err != nil {
		logger.Error(fmt.Sprintf("Error retrieving databases: %s", err.Error()))
		return nil, err
	}
	defer result.Close()

	var feeds []types.RSSFeed
	for result.Next() {
		feed := types.RSSBuildRSSFeedFromDBRow(result)
		feeds = append(feeds, feed)
	}

	return feeds, nil
}
