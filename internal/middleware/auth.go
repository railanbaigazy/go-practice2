package middleware

import (
	"log"
	"net/http"

	"github.com/railanbaigazy/go-practice2/internal/helpers"
)

var secret = "secret123"

type ErrorResponseDto struct {
	Error string `json:"error"`
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)

		if r.Header.Get("X-API-Key") != secret {
			helpers.WriteJsonError(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}
