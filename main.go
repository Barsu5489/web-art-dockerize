package main

import (
	"fmt"
	"log"
	"net/http"

	server "asciiweb/server"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Switch on the requested URL path
		switch r.URL.Path {
		case "/":

			if r.Method == http.MethodPost {
				server.HandlePostRequest(w, r)
			} else {
				http.ServeFile(w, r, "templates/index.html") // Serve the index.html file located in the templates directory
			}

		default:

			http.NotFound(w, r) // Respond with a 404 Not Found status for all other paths
		}
	})

	// The server runs asynchronously on port 8080 using a goroutine.
	go func() {
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Println(err)
		}
	}()
	fmt.Println("server running on port 8080")
	select {} // keep the main function running indefinitely.
}
