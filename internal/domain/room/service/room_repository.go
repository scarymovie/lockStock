package service

import (
	"context"
	"database/sql"
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
	query := `SELECT uid, name, code, createdAt FROM room` // Здесь могут быть дополнительные условия, если нужно
	rows, err := r.tx.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []room.Room
	for rows.Next() {
		var room room.Room
		if err := rows.Scan(&room.UID, &room.Name, &room.Code, &room.CreatedAt); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	// Проверяем наличие ошибок после завершения чтения строк
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return rooms, nil
}
