package controllers

import "net/http"

func DefaultHandler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello World 1234"));
}