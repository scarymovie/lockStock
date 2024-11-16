package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"lockStock/internal/domain/room"
	"lockStock/internal/domain/room/repository"
	"log"
)

type GetRoomByCodeService struct {
	db *sql.DB
}

func NewGetRoomByCodeService(db *sql.DB) *GetRoomByCodeService {
	return &GetRoomByCodeService{db: db}
}

func (s *GetRoomByCodeService) GetRoomByCode(ctx context.Context, code string) (*room.Room, error) {
	log.Println("[INFO] Starting transaction to fetch room by code")

	// Начинаем транзакцию
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		log.Printf("[ERROR] Failed to begin transaction: %v", err)
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback() // Откатываем транзакцию в случае паники
			log.Printf("[ERROR] Panic recovered: %v", p)
			panic(p) // Пробрасываем панику дальше
		}
	}()

	// Используем репозиторий для получения комнаты
	log.Println("[INFO] Initializing RoomRepository")
	repo := repository.NewRoomRepository(tx)

	log.Printf("[INFO] Fetching room by code: %s", code)
	room, err := repo.GetRoomByCode(ctx, code)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch room: %v", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			log.Printf("[ERROR] Failed to rollback transaction: %v", rollbackErr)
		} else {
			log.Println("[INFO] Transaction rolled back successfully")
		}
		return nil, fmt.Errorf("failed to fetch room: %w", err)
	}

	// Коммитим транзакцию
	log.Println("[INFO] Committing transaction")
	if err := tx.Commit(); err != nil {
		log.Printf("[ERROR] Failed to commit transaction: %v", err)
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Printf("[INFO] Successfully fetched room: %+v", room)
	return room, nil
}
