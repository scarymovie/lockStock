package http

import (
	"context"
	user "lockStock/internal/domain/user"
	"lockStock/internal/usecase/user/contract"
	"log"
	"net/http"
)

type UserHandler struct {
	userService contract.UserCreator
}

func NewUserHandler(userService contract.UserCreator) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	newUser := user.NewUser("public-unique-uid")

	if _, err := h.userService.CreateUser(context.Background(), newUser); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	log.Println("User created successfully")
	w.Write([]byte("user created!"))
}
