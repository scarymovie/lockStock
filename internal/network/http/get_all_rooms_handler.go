package http

import (
	"database/sql"
	"encoding/json"
	"lockStock/internal/usecase/room/usecase"
	"log"
	"net/http"
)

type GetAllRoomsHandler struct {
	DB                *sql.DB
	GetAllRoomService *usecase.GetAllRoomService
}

func NewGetAllRoomsHandler(db *sql.DB) *GetAllRoomsHandler {
	roomService := usecase.NewGetAllRoomService(db)
	return &GetAllRoomsHandler{GetAllRoomService: roomService}
}

func (h *GetAllRoomsHandler) GetAllActiveRooms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Получаем все комнаты через сервис
	rooms, err := h.GetAllRoomService.GetAllRooms(ctx)
	if err != nil {
		http.Error(w, "Failed to fetch rooms", http.StatusInternalServerError)
		log.Printf("Error fetching rooms: %v", err)
		return
	}

	// Возвращаем успешный ответ с данными (например, в JSON формате)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(rooms); err != nil {
		http.Error(w, "Failed to encode rooms", http.StatusInternalServerError)
		log.Printf("Error encoding rooms: %v", err)
	}
}
