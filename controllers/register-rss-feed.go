package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/michaelrbond/go-rss-aggregator/api-responses"
	"github.com/michaelrbond/go-rss-aggregator/types"
)

// RegisterRssFeed registers a RSS feed
func RegisterRssFeed(res http.ResponseWriter, req *http.Request, context *types.Context) {
	decoder := json.NewDecoder(req.Body)
	var feed types.RSSFeedBase
	err := decoder.Decode(&feed)

	if err != nil && err.Error() == "EOF" {
		response := apiResponses.ErrorBadRequest("Post body not provided.")
		apiResponses.Send(response, res)
		return
	} else if err != nil {
		response := apiResponses.ErrorInternalServer()
		apiResponses.Send(response, res)
		return
	}

	if err := feed.Verify(); err != nil {
		response := apiResponses.ErrorBadRequest(err.Error())
		apiResponses.Send(response, res)
		return
	}

	if err := feed.Save(context); err != nil {
		response := apiResponses.ErrorInternalServer()
		apiResponses.Send(response, res)
		return
	}

	response := apiResponses.OkMsg("Successfully added RSS feed.")
	apiResponses.Send(response, res)
}
