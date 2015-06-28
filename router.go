package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/coopernurse/gorp"
	"github.com/gorilla/mux"
)

func FileServerRouteG(m *mux.Router, path, dir string) {
	m.PathPrefix(path).Handler(
		http.StripPrefix(path, http.FileServer(http.Dir(dir))))
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/index.html")
}

func makeLocationsRoute(dbMap *gorp.DbMap) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var locations []Location

		_, err := dbMap.Select(&locations, "SELECT * FROM locations")
		if err != nil {
			log.Fatal(err)
		}

		locationsJSON, err := json.Marshal(locations)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "%s", locationsJSON)
	}
}

func addStaticRoutes(m *mux.Router, pathsAndDirs ...string) {
	for i := 0; i < len(pathsAndDirs)-1; i += 2 {
		FileServerRouteG(m, pathsAndDirs[i], pathsAndDirs[i+1])
	}
}

func initRouter(dbMap *gorp.DbMap) *mux.Router {
	locationsRoute := makeLocationsRoute(dbMap)

	r := mux.NewRouter()
	addStaticRoutes(r, "/partials/", "public/partials",
		"/scripts/", "public/scripts", "/styles/", "public/styles",
		"/images/", "public/images")
	r.HandleFunc("/locations", locationsRoute)
	r.PathPrefix("/").HandlerFunc(indexRoute)

	return r
}
