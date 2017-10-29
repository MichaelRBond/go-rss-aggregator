package types

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

// TODO : Extend gofeed's Item
func rssItemTime(item *gofeed.Item) int32 {
	if item.UpdatedParsed != nil {
		fmt.Printf("Updated: %+v\n", item.UpdatedParsed)
	}
	if item.PublishedParsed != nil {
		fmt.Printf("Published: %+v\n", item.PublishedParsed)
	}

	return -1
}
