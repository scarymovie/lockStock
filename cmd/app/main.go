package main

import (
	"lockStock/internal/infrastructure/db"
	"lockStock/internal/middleware"
	appRouter "lockStock/internal/router"
	"lockStock/pkg/logger"
	"log"
	"net/http"
)

func main() {
	logger.Logger.Println("Starting application...")

	// Инициализация подключения к базе данных
	database, err := db.NewDBConnection()
	if err != nil {
		logger.Logger.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.Close()

	// Настройка маршрутизатора
	router := http.NewServeMux()
	appRouter.LoadRoutes(router, database)
	router.HandleFunc("/", handleOther)

	// Настройка сервера
	stack := middleware.CreateStack(middleware.Logging)
	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	logger.Logger.Println("Starting server...")
	if err := server.ListenAndServe(); err != nil {
		logger.Logger.Fatalf("Server failed to start: %v", err)
	}
}

func handleOther(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a non-domain request")
	w.Write([]byte("Hello, stranger..."))
}
