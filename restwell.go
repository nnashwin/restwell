package restwell

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Routes struct {
	Routes []Route `json:"routes"`
}

type Route struct {
	Path    string `json:"path"`
	Payload string `json:"payload"`
}

type RouteHandler struct {
	jsonStr string `json:"payload"`
}

func (rh *RouteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(rh.jsonStr))
}

func CreateMuxFromJSON(jsonStr string) *http.ServeMux {
	var routes Routes
	mux := http.NewServeMux()

	err := json.Unmarshal([]byte(jsonStr), &routes)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	for _, r := range routes.Routes {
		mux.Handle("/"+r.Path, &RouteHandler{jsonStr: r.Payload})
	}

	return mux
}
