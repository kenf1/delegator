package routes

import (
	"fmt"
	"net/http"
)

// HandleEntry
//
//	@Summary		Delegator entrypoint
//	@Description	Simple health-check or greeting endpoint.
//	@Tags			General
//	@Produce		plain
//	@Success		200	{string}	string	"Delegator entrypoint"
//	@Failure		500	{string}	string	"Internal server error"
//	@Router			/ [get]
func HandleEntry(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Delegator entrypoint")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
