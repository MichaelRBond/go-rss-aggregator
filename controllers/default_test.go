package controllers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultHandler(t *testing.T) {
	assert := assert.New(t)

	resp := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	DefaultHandler(resp, req, nil)

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
	}
	resultStr := string(result)
	assert.Equal(resp.Code, 200)
	assert.NotContains(resultStr, "error")
	assert.Contains(resultStr, "default handler")
}
