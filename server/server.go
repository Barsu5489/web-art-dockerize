package asciiweb

import (
	"fmt"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome")
}