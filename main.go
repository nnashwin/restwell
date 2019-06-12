package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const jsonStr = `{"routes":[{"route":"cookies","payload":"chocolateChip"},{"route":"snacks","payload":"{\"cookies\":\"vanilla\",\"cupcakeTypes\":[\"happiness\",\"chocolateChip\"]}"}]}`

type Routes struct {
	Routes []Route `json:"routes"`
}

type Route struct {
	Route   string `json:"route"`
	Payload string `json:"payload"`
}

func handler(str string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Serving up %s", str)
	}
}

func main() {
	var routes Routes

	err := json.Unmarshal([]byte(jsonStr), &routes)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(routes.Routes)

	for _, r := range routes.Routes {
		http.HandleFunc("/"+r.Route, handler(r.Payload))
	}

	log.Fatal(http.ListenAndServe(":8081", nil))
}
