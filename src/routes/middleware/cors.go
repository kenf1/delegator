package middleware

import (
	"net/http"
	"strings"

	"github.com/kenf1/delegator/src/models"
)

func createAllowedRequests(allowedMethods string) map[string]struct{} {
	allowedRequests := make(map[string]struct{})
	for _, method := range strings.Split(allowedMethods, ",") {
		m := strings.TrimSpace(method)
		if m != "" {
			allowedRequests[m] = struct{}{}
		}
	}

	return allowedRequests
}

func DefaultCorsMiddleware(
	next http.Handler,
	serverConfig models.ServerAddr,
	allowedMethods string,
) http.Handler {
	allowedRequests := createAllowedRequests(allowedMethods)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowedOrigin := serverConfig.Host + ":" + serverConfig.Port
		origin := r.Header.Get("Origin")

		//origin must match allowed options
		if origin != "" && origin != allowedOrigin {
			http.Error(w, "Forbidden - invalid origin", http.StatusForbidden)
			return
		}

		if r.Method != http.MethodOptions {
			if _, ok := allowedRequests[r.Method]; !ok {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}
		}

		w.Header().Set("Access-Control-Allow-Methods", allowedMethods)
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if origin == allowedOrigin {
			w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
			w.Header().Set("Vary", "Origin")
		}

		//method must match allowed options
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		//pass control to next handler
		next.ServeHTTP(w, r)
	})
}
