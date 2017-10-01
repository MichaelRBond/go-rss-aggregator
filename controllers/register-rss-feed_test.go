package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_verifyPayload(t *testing.T) {
	assert := assert.New(t)

	payloadValid := apiRegisterPayload{
		Title: "My Title",
		URL:   "http://abc.com",
	}

	payloadNoTitle := apiRegisterPayload{
		URL: "http://abc.com",
	}

	payloadNoURL := apiRegisterPayload{
		Title: "My Title",
	}

	payloadNoTitleURL := apiRegisterPayload{}

	assert.Nil(verifyPayload(payloadValid), "Valid Payload")

	err := verifyPayload(payloadNoTitle)
	assert.Contains(err.Error(), "Feed title missing")
	assert.NotContains(err.Error(), "URL")

	err = verifyPayload(payloadNoURL)
	assert.Contains(err.Error(), "Feed URL missing")
	assert.NotContains(err.Error(), "title")

	err = verifyPayload(payloadNoTitleURL)
	assert.Contains(err.Error(), "Feed URL missing")
	assert.Contains(err.Error(), "Feed title missing")

}

func Test_RegisterRssFeed_ValidInput(t *testing.T) {
	assert := assert.New(t)

	payloadValid := `{"title": "My Title", "url": "http://abc.com"}`
	payload := strings.NewReader(payloadValid)

	resp := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/v1/feeds/add", payload)
	if err != nil {
		t.Fatal(err)
	}

	RegisterRssFeed(resp, req)
	assert.Equal(resp.Code, 200)
}

func Test_RegisterRssFeed_NoInput(t *testing.T) {
	assert := assert.New(t)
	resp := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/api/v1/feeds/add", strings.NewReader(""))
	if err != nil {
		t.Fatal(err)
	}

	RegisterRssFeed(resp, req)
	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
	}
	resultStr := string(result)
	assert.Equal(resp.Code, 400)
	assert.Contains(resultStr, "Post body not provided")
}
