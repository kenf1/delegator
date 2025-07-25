package routes

import (
	"encoding/json"
	"net/http"
	"os"

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

		response := map[string]string{"token": encodeResult}
		if err := json.NewEncoder(w).Encode(response); err != nil {
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

		response := map[string]interface{}{
			"claims": decodeResult,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func AuthRoutes(authConfig models.AuthConfig) *http.ServeMux {
	authMux := http.NewServeMux()

	authMux.HandleFunc("POST /create", GenerateJWT(authConfig))
	//todo: consider cors
	if os.Getenv("DEPLOY_STATUS") == "dev" {
		authMux.HandleFunc("GET /uncreate/{token}", DeconstructJWT(authConfig))
	}

	return authMux
}
