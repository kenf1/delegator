package routes

import (
	"encoding/json"
	"net/http"

	"github.com/kenf1/delegator/src/auth"
	"github.com/kenf1/delegator/src/models"
)

func GenerateJWT(authConfig models.AuthConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var userInfo models.UserInfo

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&userInfo); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}

		encodeResult, err := auth.EncodeJWT(userInfo, authConfig)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(encodeResult); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func DeconstructJWT(authConfig models.AuthConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		inputToken := r.PathValue("token")

		decodeResult, err := auth.DecodeJWT(inputToken, authConfig)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(decodeResult); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
