package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/railanbaigazy/go-practice2/internal/helpers"
)

type getUserResponseDto struct {
	UserId int `json:"user_id"`
}

type postUserRequestDto struct {
	Name string `json:"name"`
}

type postUserResponseDto struct {
	Created string `json:"created"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		idStr := r.URL.Query().Get("id")

		if idStr == "" {
			helpers.WriteJsonError(w, "not implemented yet", http.StatusNotFound)
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			helpers.WriteJsonError(w, "invalid id", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(getUserResponseDto{UserId: id})
	case http.MethodPost:
		var v postUserRequestDto

		if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
			helpers.WriteJsonError(w, err.Error(), http.StatusBadRequest)
			return
		}

		if v.Name == "" {
			helpers.WriteJsonError(w, "invalid name", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(postUserResponseDto{Created: v.Name})

	default:
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
	}

}
