package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const testResponse1 = "Hello world"
const testResponse2 = "Hello, butts"

func makeTestServer(response string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, response)
	}))
}

func TestItFetchesURLs(t *testing.T) {
	testServer1 := makeTestServer(testResponse1)
	defer testServer1.Close()

	testServer2 := makeTestServer(testResponse2)
	defer testServer2.Close()

	responses := &bytes.Buffer{}
	urls := []string{testServer1.URL, testServer2.URL}

	expectedResponse1report := fmt.Sprintf(sizeFormat, len(testResponse1), testServer1.URL)
	expectedResponse2report := fmt.Sprintf(sizeFormat, len(testResponse2), testServer2.URL)

	fetchFast(urls, responses)

	if !strings.Contains(responses.String(), expectedResponse1report) {
		t.Error("Expected to see report", expectedResponse1report, "in", responses.String())
	}

	if !strings.Contains(responses.String(), expectedResponse2report) {
		t.Error("Expected to see report", expectedResponse2report, "in", responses.String())
	}
}
