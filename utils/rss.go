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

// RSSGetAllGroups retrieves all the groups from the database, returns an array
func RSSGetAllGroups(context *types.Context) ([]types.RSSGroup, error) {
	result, err := context.Db.Query("SELECT * FROM `groups`")
	if err != nil {
		logger.Error(fmt.Sprintf("Error retrieving groups: %s", err.Error()))
		return nil, err
	}
	defer result.Close()

	var groups []types.RSSGroup
	for result.Next() {
		group, err := types.RSSBuildRSSGroupFromDBRow(result)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}
