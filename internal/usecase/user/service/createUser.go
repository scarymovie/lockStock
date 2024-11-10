package service

import (
	"context"
	"database/sql"
	"lockStock/internal/domain/user"
	"lockStock/internal/domain/user/service"
)

// UserService предоставляет методы для работы с пользователями.
type UserService struct {
	db *sql.DB
}

// NewUserService создаёт новый сервис с подключением к базе данных.
func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

// CreateUser создаёт нового пользователя, используя транзакцию.
func (s *UserService) CreateUser(ctx context.Context, newUser *user.User) (string, error) {
	// Начинаем транзакцию
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return "", err
	}

	// Выполняем сохранение пользователя
	repo := service.NewUserRepository(tx)
	id, err := repo.SaveUser(ctx, newUser)
	if err != nil {
		tx.Rollback() // Откатываем транзакцию в случае ошибки
		return "", err
	}

	// Фиксируем транзакцию
	if err := tx.Commit(); err != nil {
		return "", err
	}

	return id, nil
}
