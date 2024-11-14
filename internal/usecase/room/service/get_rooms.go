package service

import (
	"context"
	"database/sql"
	"lockStock/internal/domain/room"
	"lockStock/internal/domain/room/service"
)

type RoomService struct {
	db *sql.DB
}

func NewRoomService(db *sql.DB) *RoomService {
	return &RoomService{db: db}
}

// Получаем все комнаты через репозиторий
func (s *RoomService) GetAllRooms(ctx context.Context) ([]room.Room, error) {
	// Начинаем транзакцию
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Используем репозиторий для получения комнат
	repo := service.NewRoomRepository(tx)
	rooms, err := repo.GetAllRooms(ctx)
	if err != nil {
		tx.Rollback() // Откатываем транзакцию в случае ошибки
		return nil, err
	}

	// Коммитим транзакцию, так как запрос успешен
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return rooms, nil
}
