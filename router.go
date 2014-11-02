package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/coopernurse/gorp"
	"github.com/gorilla/mux"
)

func indexRoute(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "views/index.html")
}

func makeLocationsRoute(dbMap *gorp.DbMap) func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		var locations []Location
		var locationsJSON []byte

		_, err := dbMap.Select(&locations, "SELECT * FROM locations")

		if err != nil {
			log.Fatal(err)
		}

		locationsJSON, err = json.Marshal(locations)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "%s", locationsJSON)
	}
}

func initRouter(dbMap *gorp.DbMap) *mux.Router{
	locationsRoute := makeLocationsRoute(dbMap)

	r := mux.NewRouter()
	r.HandleFunc("/locations", locationsRoute)
	r.HandleFunc("/{url:.*}", indexRoute)

	return r
}
