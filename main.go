package main

import (
	"fmt"
	"html/template"
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

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

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
				// Serve the index.html file located in the templates directory
				tpl, _ := template.ParseFiles("templates/index.html")

				tpl.Execute(w, nil)
			}

		default:
			// Respond with a 404 Not Found status for all other paths
			tpl, _ := template.ParseFiles("templates/404.html")
			w.WriteHeader(http.StatusNotFound)

			tpl.Execute(w, nil)

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
