package apiResponses

import (
	"encoding/json"
)

// APIResponse is a common interface for API Responses
type APIResponse interface {
	getStatusCode() int
	toJSON() []byte
}

// APIResponseDefault is a standard response
type APIResponseDefault struct {
	Meta APIResponseMeta
	Data APIResponseMessage
}

// APIResponseDefaultError is a standard response with errors
type APIResponseDefaultError struct {
	Meta APIResponseMetaWithError
	Data APIResponseMessage
}

// APIResponseMeta is the metadata that goes to standard responses from the API
type APIResponseMeta struct {
	StatusCode int
}

// APIResponseMetaWithError is metadata with error information
type APIResponseMetaWithError struct {
	StatusCode int
	Error      APIResponseError
}

// APIResponseMessage is a basic data struct that contains only a single message
type APIResponseMessage struct {
	Message string
}

// APIResponseError is metadata about any errors that occured
type APIResponseError struct {
	ID          string
	Description string
}

func buildAPIResponseMeta(statusCode int) APIResponseMeta {
	return APIResponseMeta{
		StatusCode: statusCode,
	}
}

func buildAPIResponseMetaWithError(statusCode int, errorID string, errorDescription string) APIResponseMetaWithError {
	return APIResponseMetaWithError{
		StatusCode: statusCode,
		Error: APIResponseError{
			ID:          errorID,
			Description: errorDescription,
		},
	}
}

// ToJSON converts a APIResponseDefault to json
func (api APIResponseDefault) toJSON() []byte {
	responseBytes, _ := json.Marshal(api)
	return responseBytes
}

func (api APIResponseDefault) getStatusCode() int {
	return api.Meta.StatusCode
}

// ToJSON converts a APIResponseDefaultError to json
func (api APIResponseDefaultError) toJSON() []byte {
	responseBytes, _ := json.Marshal(api)
	return responseBytes
}

func (api APIResponseDefaultError) getStatusCode() int {
	return api.Meta.StatusCode
}
