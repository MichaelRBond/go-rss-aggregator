package types

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
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

// Verify validates an RSSFeedBase
func (feed RSSFeedBase) Verify() error {
	var errs []string
	if feed.Title == "" {
		errs = append(errs, "Feed title missing")
	}
	if feed.URL == "" {
		errs = append(errs, "Feed URL missing")
	}

	if len(errs) == 0 {
		return nil
	}

	errorMsg := strings.Join(errs, ", ")
	return errors.New(errorMsg)
}

// RSSFeed describes an RSS feed in saved in the database
type RSSFeed struct {
	RSSFeedBase
	ID          int16
	LastUpdated int32
}

// Save saves an RSSFeed to the database
func (feed RSSFeed) Save(context *Context) error {
	_, err := context.Db.Query("UPDATE `feeds` SET `title`=?, `url`=?, `lastUpdated`=? WHERE `id`=?",
		feed.Title, feed.URL, feed.LastUpdated, feed.ID)
	if err != nil {
		logger.Error(fmt.Sprintf("Error, %s, while updating feed id=%d", err.Error(), feed.ID))
	}
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

	for _, item := range feedContent.Items {
		itemTime := rssItemTime(item)
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

	itemGUID := ""
	if item.GUID != "" {
		itemGUID = item.GUID
	}
	guid := fmt.Sprintf("%s:%s", itemGUID, item.Link)

	base := RSSItemBase{}
	base.FeedID = feed.ID
	base.Title = item.Title
	base.Description = item.Description
	base.Content = item.Content
	base.Link = item.Link
	if item.UpdatedParsed != nil {
		base.Updated = int32(item.UpdatedParsed.Unix())
	} else {
		base.Updated = 0
	}
	base.Published = int32(item.PublishedParsed.Unix())
	if item.Author != nil {
		base.Author = item.Author.Name
	} else {
		base.Author = ""
	}
	base.GUID = guid
	if item.Image != nil {
		base.Image = item.Image.URL
	} else {
		base.Image = ""
	}
	base.Categories = make([]string, 1, 1) // TODO :
	base.Enclosures = ""                   // TODO :
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
func (item RSSItemBase) Save(context *Context) error {
	if rssItem, _ := GetRSSItemByGUID(context, item.GUID); rssItem.ID != 0 {
		item.Update(context)
	}
	// TODO : Save Categories to the database
	_, err := context.Db.Query("INSERT INTO `feedItems` (`feedId`, `title`, `description`, `content`, `link`, "+
		"`updated`, `published`, `author`, `guid`, `image`, `enclosures`) "+
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		item.FeedID, item.Title, item.Description, item.Content, item.Link, item.Updated, item.Published,
		item.Author, item.GUID, item.Image, item.Enclosures)
	if err != nil {
		logger.Error(fmt.Sprintf("Error, %s, Saving RSSItemBase=%v", err.Error(), item))
		return err
	}
	return nil
}

// Update updates an RSSItemBase in the database
func (item RSSItemBase) Update(context *Context) error {
	_, err := context.Db.Query("UPDATE `feedItems` SET `title`=?, `description`=?, `content`=?, `link`=?, "+
		"`updated`=?, `published`=?, `author`=?, `image`=?, `categories`=?, `enclosures`=?) WHERE `guid`=?"+
		"VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		item.Title, item.Description, item.Content, item.Link, item.Updated, item.Published,
		item.Author, item.Image, item.Categories, item.Enclosures, item.GUID)
	if err != nil {
		logger.Error(fmt.Sprintf("Error Updating RSSItemBase=%v", item))
		return err
	}
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

// RSSGroupBase base group struct
type RSSGroupBase struct {
	Name string
}

// Save saves an RSSGroupBase to the database
func (group RSSGroupBase) Save(context *Context) error {
	_, err := context.Db.Query("INSERT INTO `groups` (`name`) VALUES(?);", group.Name)
	if err != nil {
		logger.Error(fmt.Sprintf("Error saving new group: %s", err.Error()))
		return err
	}
	return nil
}

// Verify ensures that group s valid
func (group RSSGroupBase) Verify() error {
	var errs []string
	if group.Name == "" {
		errs = append(errs, "Group name missing")
	}

	if len(errs) == 0 {
		return nil
	}

	errorMsg := strings.Join(errs, ", ")
	return errors.New(errorMsg)
}

// RSSGroup as saved in the database
type RSSGroup struct {
	RSSGroupBase
	ID int32
}

// Save updates an RSSGroup in the database
func (group RSSGroup) Save(context *Context) error {
	_, err := context.Db.Query("UPDATE `groups` SET `name`=? WHERE `id`=?", group.Name, group.ID)
	if err != nil {
		logger.Error(fmt.Sprintf("Error updating group: %s", err.Error()))
		return err
	}
	return nil
}

// RSSGroupAdd adds a feed to a group
type RSSGroupAdd struct {
	FeedID  int32 `json:"feed_id"`
	GroupID int32 `json:"group_id"`
}

// Save saves a feed/group association to the database
func (groupAdd RSSGroupAdd) Save(context *Context) error {
	// TODO : Verify that group id and feed id are valid
	_, err := context.Db.Query("INSERT INTO `feedGroups` (`feedId`, `groupId`) VALUES(?, ?)", groupAdd.FeedID, groupAdd.GroupID)
	if err != nil {
		logger.Error(fmt.Sprintf("Error adding feed/group association: %s", err.Error()))
	}
	return nil
}

// Verify validates an RSSGroupAdd
func (groupAdd RSSGroupAdd) Verify() error {
	fmt.Printf("groupAdd: %+v\n", groupAdd)
	var errs []string
	if groupAdd.FeedID == 0 {
		errs = append(errs, "`feed_id` missing")
	}
	if groupAdd.GroupID == 0 {
		errs = append(errs, "`group_id` missing")
	}

	if len(errs) == 0 {
		return nil
	}

	errorMsg := strings.Join(errs, ", ")
	return errors.New(errorMsg)
}
