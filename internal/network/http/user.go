package http

import (
	"database/sql"
	"github.com/google/uuid"
	"lockStock/internal/domain/user"
	"lockStock/internal/usecase/user/usecase"
	"lockStock/pkg/logger"
	"net/http"
)

type UserHandler struct {
	DB          *sql.DB
	UserService *usecase.UserService // Указываем полный путь к структуре UserService
}

// NewUserHandler конструктор для создания UserHandler и инициализации UserService
func NewUserHandler(db *sql.DB) *UserHandler {
	userService := usecase.NewUserService(db) // Имя переменной с маленькой буквы
	return &UserHandler{DB: db, UserService: userService}
}

// CreateUser обрабатывает создание пользователя с использованием транзакции
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	logger.Logger.Println("Starting to create a new user transaction...")

	// Запуск транзакции
	ctx := r.Context()
	tx, err := h.DB.BeginTx(ctx, nil)
	if err != nil {
		logger.Logger.Println("Transaction failed:", err)
		http.Error(w, "Failed to start transaction", http.StatusInternalServerError)
		return
	}
	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				logger.Logger.Printf("Failed to rollback transaction: %v", rbErr)
			}
		} else {
			if cmErr := tx.Commit(); cmErr != nil {
				logger.Logger.Printf("Failed to commit transaction: %v", cmErr)
				http.Error(w, "Failed to commit transaction", http.StatusInternalServerError)
			}
		}
	}()

	// Создание пользователя
	newUser := user.NewUser("u" + uuid.New().String())
	if _, err = h.UserService.CreateUser(ctx, newUser); err != nil {
		logger.Logger.Printf("Error creating user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created: " + newUser.UID))
}
