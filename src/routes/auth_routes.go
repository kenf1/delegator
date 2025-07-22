package routes

import (
	"encoding/json"
	"net/http"

	"github.com/kenf1/delegator/src/auth"
	"github.com/kenf1/delegator/src/models"
)

func GenerateJWT(w http.ResponseWriter, r *http.Request) {
	input := r.PathValue("value")

	userInfo := models.UserInfo{
		Id:          input,
		Email:       "blah",
		Roles:       []string{"admin"},
		Permissions: []string{"read", "write", "delete"},
		Org_id:      69,
	}

	authConfig := models.AuthConfig{
		SecretKey: []byte("42069"),
		Issuer:    "me",
	}

	res, err := auth.EncodeJWT(userInfo, authConfig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeconstructJWT(w http.ResponseWriter, r *http.Request) {
	input := r.PathValue("value")

	authConfig := models.AuthConfig{
		SecretKey: []byte("42069"),
		Issuer:    "me",
	}

	res, err := auth.DecodeJWT(input, authConfig.SecretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
