package controllers

import (
	"encoding/json"
	"net/http"
)

// APIResponseDefault is a standard response
type APIResponseDefault struct {
	Meta APIResponseMeta
	Data APIResponseMessage
}

// APIResponseMeta is the metadata that goes to standard responses from the API
type APIResponseMeta struct {
	StatusCode int
}

// APIResponseMessage is a basic data struct that contains only a single message
type APIResponseMessage struct {
	Message string
}

// OkMsg returns a a 200 status code and string message as data
func OkMsg(message string) *APIResponseDefault {
	return &APIResponseDefault{
		Meta: APIResponseMeta{200},
		Data: APIResponseMessage{message},
	}
}

// DefaultHandler renders the hompage
func DefaultHandler(res http.ResponseWriter, req *http.Request) {
	response := OkMsg("default handler 1234")
	responseBytes, _ := json.Marshal(response)
	res.Write(responseBytes)
}
