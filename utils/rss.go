package utils

import (
	"database/sql"
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
		feed := rssBuildRSSFeedFromDBRow(result)
		feeds = append(feeds, feed)
	}

	return feeds, nil
}

func rssBuildRSSFeedFromDBRow(row *sql.Rows) types.RSSFeed {
	var feed types.RSSFeed
	err := row.Scan(&feed.ID, &feed.Title, &feed.URL, &feed.LastUpdated)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating feed struct: %s", err.Error()))
		// TODO : Better error handling
	}
	return feed
}

func rssBuildRSSItemFromDBRow(row *sql.Rows) types.RSSItem {
	var item types.RSSItem
	err := row.Scan(&item.ID, &item.FeedID, &item.Title, &item.Description, &item.Content, &item.Link, &item.Updated,
		&item.Published, &item.Author, &item.GUID, &item.Image, &item.Categories, &item.Enclosures, &item.Read,
		&item.Starred)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating RSS Item struct: %s", err.Error()))
	}
	return item
}
