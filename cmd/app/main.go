// main.go
package main

import (
	"database/sql"
	"lockStock/internal/middleware"
	appRouter "lockStock/internal/router"
	"lockStock/pkg/logger"
	"log"
	"net/http"

	_ "github.com/lib/pq" // PostgreSQL драйвер
)

func main() {
	logger.Logger.Println("Starting application...")

	// Подключение к базе данных
	db, err := sql.Open("postgres", "host=dev-db port=5432 user=db_user password=db_password dbname=db_database sslmode=disable")
	if err != nil {
		logger.Logger.Fatalf("Error opening database: %v", err.Error())
	}
	defer db.Close()

	// Проверка подключения к базе данных
	if err := db.Ping(); err != nil {
		logger.Logger.Fatalf("Database not reachable: %v", err.Error())
	}

	// Создаем маршрутизатор и передаем в него подключение к базе данных
	router := http.NewServeMux()
	appRouter.LoadRoutes(router, db)
	router.HandleFunc("/", handleOther)

	// Создаем стек промежуточных обработчиков
	stack := middleware.CreateStack(
		middleware.Logging,
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: stack(router),
	}

	logger.Logger.Println("Starting server...")
	if err := server.ListenAndServe(); err != nil {
		logger.Logger.Fatalf("Server failed to start: %v", err.Error())
	}
}

func handleOther(w http.ResponseWriter, r *http.Request) {
	log.Println("Received a non-domain request")
	w.Write([]byte("Hello, stranger..."))
}
