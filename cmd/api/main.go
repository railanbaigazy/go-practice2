package main

import (
	"fmt"
	"net/http"

	"github.com/railanbaigazy/go-practice2/internal/handlers"
	"github.com/railanbaigazy/go-practice2/internal/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/user", middleware.AuthMiddleware(handlers.UserHandler))

	fmt.Println("Serving on port 8080")
	http.ListenAndServe(":8080", mux)
}
