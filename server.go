package main

import (
	"fmt"
	"net/http"
)

func addStaticRoutes() {
	http.Handle("/partials/", http.StripPrefix("/partials/",
		http.FileServer(http.Dir("public/partials"))))
	http.Handle("/scripts/", http.StripPrefix("/scripts/",
		http.FileServer(http.Dir("public/scripts"))))
	http.Handle("/styles/", http.StripPrefix("/styles/",
		http.FileServer(http.Dir("public/styles"))))
	http.Handle("/images/", http.StripPrefix("/images/",
		http.FileServer(http.Dir("public/images"))))
}

func main() {
	fmt.Println("Initializing database")
	db := initDB()

	router := initRouter(db)
	http.Handle("/", router)
	addStaticRoutes()

	fmt.Println("Starting server")
	http.ListenAndServe(":1123", nil)
}
