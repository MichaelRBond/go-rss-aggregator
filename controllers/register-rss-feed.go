package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/michaelrbond/go-rss-aggregator/api-responses"
)

type apiRegisterPayload struct {
	Title string
	URL   string
}

// RegisterRssFeed registers a RSS feed
func RegisterRssFeed(res http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var feed apiRegisterPayload
	err := decoder.Decode(&feed)

	if err != nil {
		response := apiResponses.ErrorInternalServer()
		responseBytes, _ := json.Marshal(response)
		res.WriteHeader(response.Meta.StatusCode)
		res.Write(responseBytes)
		return
	}

	// TODO : save payload to database
	// if err

	response := apiResponses.OkMsg("Successfully added RSS feed.")
	responseBytes, _ := json.Marshal(response)
	res.Write(responseBytes)
}
