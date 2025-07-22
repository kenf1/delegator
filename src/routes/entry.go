package routes

import (
	"fmt"
	"net/http"
)

func HandleEntry(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Delegator entrypoint")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
