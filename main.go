package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const jsonStr = `{"routes":[{"route":"cookies","payload":"chocolateChip"},{"route":"snacks","payload":"{\"cookies\":\"vanilla\",\"cupcakeTypes\":[\"happiness\",\"chocolateChip\"]}"}]}`

const addr = "localhost:12345"

type Routes struct {
	Routes []Route `json:"routes"`
}

type Route struct {
	Route   string `json:"route"`
	Payload string `json:"payload"`
}

type RouteHandler struct {
	jsonStr string `json:"payload"`
}

func (rh *RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, rh.jsonStr)
}

func main() {
	var routes Routes
	mux := http.NewServeMux()

	err := json.Unmarshal([]byte(jsonStr), &routes)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(routes.Routes)

	for _, r := range routes.Routes {
		mux.Handle("/"+r.Route, &RouteHandler{jsonStr: r.Payload})
	}
	log.Printf("Now Listening on %s...\n", addr)

	server := http.Server{Handler: mux, Addr: addr}

	log.Fatal(server.ListenAndServe())
}
