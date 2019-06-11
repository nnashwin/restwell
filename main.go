package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(str string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Serving up %s", str)
	}
}

func main() {
	myMap := make(map[string]string)
	myMap["cookies"] = "chocolateChip"
	myMap["cake"] = "bavarian"
	myMap["animal"] = "hamster"
	for k, v := range myMap {
		http.HandleFunc("/"+k, handler(v))
	}
	log.Fatal(http.ListenAndServe(":8081", nil))
}
