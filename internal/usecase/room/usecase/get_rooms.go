package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"lockStock/internal/domain/room"
	"lockStock/internal/domain/room/repository"
	"log"
)

type GetAllRoomService struct {
	db *sql.DB
}

func NewGetAllRoomService(db *sql.DB) *GetAllRoomService {
	return &GetAllRoomService{db: db}
}

func (s *GetAllRoomService) GetAllRooms(ctx context.Context) ([]room.Room, error) {
	log.Println("[INFO] Starting transaction to fetch all rooms")

	// Начинаем транзакцию
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("[ERROR] Failed to begin transaction: %v", err)
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Используем репозиторий для получения комнат
	log.Println("[INFO] Initializing RoomRepository")
	repo := repository.NewRoomRepository(tx)

	log.Println("[INFO] Fetching all rooms from repository")
	rooms, err := repo.GetAllRooms(ctx)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch rooms: %v", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Printf("[ERROR] Failed to rollback transaction: %v", rollbackErr)
		} else {
			log.Println("[INFO] Transaction rolled back successfully")
		}
		return nil, fmt.Errorf("failed to fetch rooms: %w", err)
	}

	// Коммитим транзакцию
	log.Println("[INFO] Committing transaction")
	if err := tx.Commit(); err != nil {
		log.Printf("[ERROR] Failed to commit transaction: %v", err)
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Printf("[INFO] Successfully fetched %d rooms", len(rooms))
	return rooms, nil
}
