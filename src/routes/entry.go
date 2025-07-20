package routes

import (
	"fmt"
	"net/http"
)

func HandleEntry(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Delegator entrypoint")
}
