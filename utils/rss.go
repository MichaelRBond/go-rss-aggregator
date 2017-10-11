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

// SaveRSSFeedToDB saves a types.RSSFeed to the database
func SaveRSSFeedToDB(context *types.Context, feed types.RSSFeed) error {

	// TODO : check for duplicates - Need to set field to unique in the database
	_, err := context.Db.Query("INSERT INTO `feeds` (`title`, `url`) VALUES(?, ?);", feed.Title, feed.URL)
	if err != nil {
		logger.Error(fmt.Sprintf("Error saving new feed: %s", err.Error()))
		return err
	}

	return nil
}

func rssBuildRSSFeedFromDBRow(row *sql.Rows) types.RSSFeed {
	var feed types.RSSFeed
	err := row.Scan(&feed.ID, &feed.Title, &feed.URL)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating feed struct: %s", err.Error()))
		// TODO : Better error handling
	}
	return feed
}
