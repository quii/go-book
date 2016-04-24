package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

const testResponseString = "Hello, world"

var (
	testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, testResponseString)
	}))
)

func TestItCanStreamURLs(t *testing.T) {
	buffer := &bytes.Buffer{}
	streamURL(testServer.URL, buffer)

	if !strings.Contains(buffer.String(), testResponseString) {
		t.Error("Server response was not written to buffer correctly, got ", buffer.String())
	}

	if !strings.Contains(buffer.String(), "200"){
		t.Error("Expected the status code 200 to be included in the output, got", buffer.String())
	}
}

func TestItPreappendsHTTPWhenNeeded(t *testing.T) {
	fixedURL := urlFixer("http://google.com")
	if fixedURL != "http://google.com" {
		t.Error("Expect already correct urls to be the same but got", fixedURL)
	}

	fixedURL = urlFixer("google.com")
	if fixedURL != "http://google.com"{
		t.Error("Expect URL to have http:// added but got", fixedURL)
	}
}
