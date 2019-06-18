package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var routeArr = [...]Route{Route{Route: "cookies", Payload: "chocolateChip"}, Route{Route: "snacks", Payload: "{\"cookies\":\"vanilla\",\"cupcakeTypes\":[\"happiness\",\"chocolateChip\"]}"}}

func TestRouteHandlers(t *testing.T) {
	for _, route := range routeArr {
		func(route Route) {
			handler := &RouteHandler{jsonStr: route.Payload}
			server := httptest.NewServer(handler)
			defer server.Close()
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal(err)
			}

			if resp.StatusCode != 200 {
				t.Fatalf("Received non-200 response: %d\n", resp.StatusCode)
			}

			expected := route.Payload
			actual, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatal(err)
			}

			if expected != string(actual) {
				t.Errorf("Expected the message '%s'\n", expected)
			}
		}(route)
	}
}
