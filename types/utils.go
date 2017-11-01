package types

import (
	"database/sql"
	"fmt"

	"github.com/michaelrbond/go-rss-aggregator/logger"
	"github.com/mmcdole/gofeed"
)

// TODO : Extend gofeed's Item
func rssItemTime(item *gofeed.Item) int32 {
	if item.UpdatedParsed != nil {
		return int32(item.UpdatedParsed.Unix())
	}
	if item.PublishedParsed != nil {
		return int32(item.PublishedParsed.Unix())
	}
	return -1
}

// GetRSSItemByGUID retieves an RSSItem from the database
func GetRSSItemByGUID(context *Context, guid string) (RSSItem, error) {
	row, err := context.Db.Query("SELECT * FROM `feedItems` WHERE `guid`=?", guid)
	if err != nil {
		return RSSItem{ID: 0}, err
	}
	var rssItem RSSItem
	for row.Next() {
		rssItem, err = rssBuildRSSItemFromDBRow(row)
	}
	return rssItem, err
}

// RSSBuildRSSFeedFromDBRow builds a RSSFeed object from a database row
func RSSBuildRSSFeedFromDBRow(row *sql.Rows) RSSFeed {
	var feed RSSFeed
	err := row.Scan(&feed.ID, &feed.Title, &feed.URL, &feed.LastUpdated)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating feed struct: %s", err.Error()))
		// TODO : Better error handling
	}
	return feed
}

func rssBuildRSSItemFromDBRow(row *sql.Rows) (RSSItem, error) {
	var item RSSItem
	err := row.Scan(&item.ID, &item.FeedID, &item.Title, &item.Description, &item.Content, &item.Link, &item.Updated,
		&item.Published, &item.Author, &item.GUID, &item.Image, &item.Categories, &item.Enclosures, &item.Read,
		&item.Starred)
	if err != nil {
		logger.Error(fmt.Sprintf("Error creating RSS Item struct: %s", err.Error()))
		return RSSItem{}, err
	}
	return item, nil
}
