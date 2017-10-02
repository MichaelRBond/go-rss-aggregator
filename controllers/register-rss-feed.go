package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/michaelrbond/go-rss-aggregator/api-responses"
)

type apiRegisterPayload struct {
	Title string
	URL   string
}

// RegisterRssFeed registers a RSS feed
func RegisterRssFeed(res http.ResponseWriter, req *http.Request, context *Context) {
	decoder := json.NewDecoder(req.Body)
	var feed apiRegisterPayload
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

	err = verifyPayload(feed)
	if err != nil {
		response := apiResponses.ErrorBadRequest(err.Error())
		apiResponses.Send(response, res)
		return
	}

	// TODO : Save to database

	response := apiResponses.OkMsg("Successfully added RSS feed.")
	apiResponses.Send(response, res)
}

func verifyPayload(payload apiRegisterPayload) error {
	var errs []string
	if payload.Title == "" {
		errs = append(errs, "Feed title missing")
	}
	if payload.URL == "" {
		errs = append(errs, "Feed URL missing")
	}

	if len(errs) == 0 {
		return nil
	}

	errorMsg := strings.Join(errs, ", ")
	return errors.New(errorMsg)
}
