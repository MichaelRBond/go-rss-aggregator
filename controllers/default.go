package controllers

import (
	"net/http"

	"github.com/michaelrbond/go-rss-aggregator/api-responses"
	"github.com/michaelrbond/go-rss-aggregator/types"
)

// DefaultHandler renders the hompage
func DefaultHandler(res http.ResponseWriter, req *http.Request, context *types.Context) {
	response := apiResponses.OkMsg("default handler 1234")
	apiResponses.Send(response, res)
}
