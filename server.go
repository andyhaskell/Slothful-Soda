package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Initializing database")
	db := initDB()

	router := initRouter(db)
	
	server := &http.Server{
		Addr:    ":1123",
		Handler: router,
	}

	fmt.Println("Starting server")
	server.ListenAndServe()
}
