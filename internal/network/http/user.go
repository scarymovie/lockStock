package http

import (
	"log"
	"net/http"
)

type UserHandler struct{}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	log.Println("received request to create a monster")
	w.Write([]byte("user created!"))
}
