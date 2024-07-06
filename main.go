package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	server "asciiweb/server"
)

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		// Handle other errors if needed
		fmt.Println("Error checking file:", err)
		return false
	}
	return true
}

func main() {
	indexFilePath := "templates/index.html"

	// Check if the index.html file exists
	if !checkFileExists(indexFilePath) {
		log.Fatalf("File %s does not exist.", indexFilePath)
		return
	}
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
