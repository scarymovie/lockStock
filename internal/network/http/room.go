package http

import (
	"database/sql"
	"encoding/json"
	"lockStock/internal/usecase/room/service"
	"log"
	"net/http"
)

type RoomHandler struct {
	DB          *sql.DB
	RoomService *service.RoomService
}

func NewRoomHandler(db *sql.DB) *RoomHandler {
	roomService := service.NewRoomService(db)
	return &RoomHandler{RoomService: roomService}
}

func (h *RoomHandler) GetAllActiveRooms(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Получаем все комнаты через сервис
	rooms, err := h.RoomService.GetAllRooms(ctx)
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
