package main

import (
	"testing"
	"net/http/httptest"
	"fmt"
	"net/http"
	"bytes"
)

const testResponseString = "Hello, world"

var(
	testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, testResponseString)
	}))
)

func TestItCanStreamURLs(t *testing.T){
	buffer := &bytes.Buffer{}
	streamURL(testServer.URL, buffer)

	if buffer.String() != testResponseString{
		t.Error("Server response was not written to buffer, got ",buffer.String())
	}
}
