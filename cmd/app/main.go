package main

import (
	"fmt"
	"lockStock/internal/middleware"
	appRouter "lockStock/internal/router"
	"log"
	"net/http"
)

const configPath = "config/main"

func main() {
	router := http.NewServeMux()
	appRouter.LoadRoutes(router)
	router.HandleFunc("/", handleOther)

	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	fmt.Println("Server listening on port :8080")
	server.ListenAndServe()
}

func handleOther(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a non domain request")
	w.Write([]byte("Hello, stranger..."))
}
