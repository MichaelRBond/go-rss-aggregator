package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/michaelrbond/go-rss-aggregator/api-responses"
	"github.com/michaelrbond/go-rss-aggregator/types"
)

// FeverAPI provides access to feeds via the fever api
func FeverAPI(res http.ResponseWriter, req *http.Request, context *types.Context) {
	decoder := json.NewDecoder(req.Body)
	var apiRequest types.FeverAPIPostRequest
	err := decoder.Decode(&apiRequest)

	if err != nil && err.Error() == "EOF" {
		response := apiResponses.ErrorBadRequest("Post body not provided.")
		apiResponses.Send(response, res)
		return
	} else if err != nil {
		response := apiResponses.ErrorInternalServer()
		apiResponses.Send(response, res)
		return
	}

	fmt.Printf("Payload: %s\n", apiRequest.APIKey)
	fmt.Printf("Query %v\n", req.URL.Query())
	fmt.Printf("Query API: %v\n", req.URL.Query().Get("api"))
	fmt.Printf("Query Groups: %v\n", req.URL.Query().Get("groups"))

	api := "json"
	feverAPIOptions := map[string]bool{
		"groups":          true,
		"feeds":           false,
		"favicons":        false,
		"items":           false,
		"links":           false,
		"unread_item_ids": false,
		"saved_item_ids":  false,
	}

	for field := range req.URL.Query() {
		// TODO : neeed to make sure we are only setting valid values
		feverAPIOptions[field] = true
		if field == "api" {
			api = req.URL.Query().Get(field)
		}
	}

	fmt.Printf("options: %v", feverAPIOptions)
	fmt.Printf("api %s", api)

	response := apiResponses.OkMsg("Done.")
	apiResponses.Send(response, res)
}
