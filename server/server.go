package asciiweb

import (
	"fmt"
	"net/http"
)

func HandlePostRequest(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welco")
}
