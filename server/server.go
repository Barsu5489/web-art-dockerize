package asciiweb

import (
	art "asciiweb/server/ascii-art"
	"html/template"
	"log"
	"net/http"
)

type Result struct {
	Data string
}

func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "400 Bad Request: Unable to parse form", http.StatusBadRequest)
		return
	}

	// Retrieve form values
	userInput := r.FormValue("text")
	bannerFile := r.FormValue("banner")

	// Generate ASCII art
	res := art.GenArt(userInput, bannerFile)

	// Prepare the result
	var artOutput Result
	if res == "500" {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		return
	} else if res == "" {
		artOutput = Result{Data: "400 Bad Request"}
		w.WriteHeader(http.StatusBadRequest)
	} else {
		artOutput = Result{Data: res}
	}

	// Parse and execute the template
	tpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error: Unable to load template", http.StatusInternalServerError)
		return
	}
	if err := tpl.Execute(w, artOutput); err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "500 Internal Server Error: Unable to execute template", http.StatusInternalServerError)
	}
}
