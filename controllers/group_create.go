package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/michaelrbond/go-rss-aggregator/api-responses"
	"github.com/michaelrbond/go-rss-aggregator/types"
)

// GroupCreate Creates a new group
func GroupCreate(res http.ResponseWriter, req *http.Request, context *types.Context) {
	decoder := json.NewDecoder(req.Body)
	var group types.RSSGroupBase
	err := decoder.Decode(&group)

	if err != nil && err.Error() == "EOF" {
		response := apiResponses.ErrorBadRequest("Post body not provided.")
		apiResponses.Send(response, res)
		return
	} else if err != nil {
		response := apiResponses.ErrorInternalServer()
		apiResponses.Send(response, res)
		return
	}

	if err := group.Verify(); err != nil {
		response := apiResponses.ErrorBadRequest(err.Error())
		apiResponses.Send(response, res)
		return
	}

	if err := group.Save(context); err != nil {
		response := apiResponses.ErrorInternalServer()
		apiResponses.Send(response, res)
		return
	}

	response := apiResponses.OkMsg("Successfully added RSS group.")
	apiResponses.Send(response, res)
}
