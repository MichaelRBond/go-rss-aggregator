package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/michaelrbond/go-rss-aggregator/api-responses"
)

// DefaultHandler renders the hompage
func DefaultHandler(res http.ResponseWriter, req *http.Request) {
	response := apiResponses.OkMsg("default handler 1234")
	responseBytes, _ := json.Marshal(response)
	res.Write(responseBytes)
}
