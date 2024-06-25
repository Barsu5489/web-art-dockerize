package asciiweb

import (
	art "asciiweb/server/ascii-art"
	// "fmt"
	"net/http"
	"html/template"
)

type Result struct {
	Data string
}
func HandlePostRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userInput := r.FormValue("text")
	bannerFile := r.FormValue("banner")

	artOutput := Result{
		Data: art.GenArt(userInput, bannerFile) ,
		}

	tpl, _:= template.ParseFiles("templates/index.html")

	tpl.Execute(w, artOutput)
}
