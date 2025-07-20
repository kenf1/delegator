package routes

import (
	"fmt"
	"log"
	"net/http"
)

func HandleEntry(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Delegator entrypoint")
	if err != nil {
		log.Printf("failed to write response: %v", err)
	}
}
