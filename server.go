package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Initializing database")
	db := initDB()

	router := initRouter(db)
	http.Handle("/", router)

	fmt.Println("Starting server")
	http.ListenAndServe(":1123", nil)
}
