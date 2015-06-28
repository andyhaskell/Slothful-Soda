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

//AddStaticRoutes takes in a Gorilla mux Router and an alternating set of URL
//paths and directory paths and for each pair of strings, the router is given a
//FileServer Handler where the first string is the URL path and the second
//string is the directory path to serve files from.
func AddStaticRoutes(m *mux.Router, pathsAndDirs ...string) {
	for i := 0; i < len(pathsAndDirs)-1; i += 2 {
		FileServerRouteG(m, pathsAndDirs[i], pathsAndDirs[i+1])
	}
}

//initRouter takes in a GORP DbMap and initializes the router's routes while
//using the DbMap to handle database functionality.
func initRouter(dbMap *gorp.DbMap) *mux.Router {
	r := mux.NewRouter()

	//Add static routes for the public directory
	AddStaticRoutes(r, "/partials/", "public/partials",
		"/scripts/", "public/scripts", "/styles/", "public/styles",
		"/images/", "public/images")
		
	//Add the locations route API with makeLocationsRoute
	r.HandleFunc("/locations", makeLocationsRoute(dbMap))
	
	//Serve all other requests with index.html, and ultimately the front-end
	//Angular.js app.
	r.PathPrefix("/").HandlerFunc(indexRoute)

	return r
}
