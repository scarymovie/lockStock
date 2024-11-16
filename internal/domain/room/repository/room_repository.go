package repository

import (
	"context"
	"database/sql"
	"fmt"
	"lockStock/internal/domain/room"
)

type RoomRepository struct {
	tx *sql.Tx
}

func NewRoomRepository(tx *sql.Tx) *RoomRepository {
	return &RoomRepository{tx: tx}
}

// Получаем все комнаты
func (r *RoomRepository) GetAllRooms(ctx context.Context) ([]room.Room, error) {
	const query = `
		SELECT uid, name, code, created_at 
		FROM rooms
	` // Здесь могут быть дополнительные условия, если нужно

	rows, err := r.tx.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query to fetch rooms: %w", err)
	}
	defer rows.Close()

	var rooms []room.Room
	for rows.Next() {
		var currentRoom room.Room
		if err := rows.Scan(&currentRoom.UID, &currentRoom.Name, &currentRoom.Code, &currentRoom.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan room row: %w", err)
		}
		rooms = append(rooms, currentRoom)
	}

	// Проверяем наличие ошибок после завершения чтения строк
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during rows iteration: %w", err)
	}

	return rooms, nil
}

func (r *RoomRepository) GetRoomByCode(ctx context.Context, code string) (*room.Room, error) {
	const query = `
		SELECT uid, name, code, created_at
		FROM rooms
		WHERE code = $1
	`

	var currentRoom room.Room
	err := r.tx.QueryRowContext(ctx, query, code).Scan(
		&currentRoom.UID,
		&currentRoom.Name,
		&currentRoom.Code,
		&currentRoom.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("room with code %q not found: %w", code, err)
		}
		return nil, fmt.Errorf("failed to fetch room by code: %w", err)
	}

	return &currentRoom, nil
}
