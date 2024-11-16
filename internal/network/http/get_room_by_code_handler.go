package http

import (
	"database/sql"
	"encoding/json"
	"lockStock/internal/usecase/room/usecase"
	"log"
	"net/http"
)

type GetRoomByCodeHandler struct {
	DB                   *sql.DB
	GetRoomByCodeService *usecase.GetRoomByCodeService
}

func NewGetRoomByCodeHandler(db *sql.DB) *GetRoomByCodeHandler {
	roomService := usecase.NewGetRoomByCodeService(db)
	return &GetRoomByCodeHandler{GetRoomByCodeService: roomService}
}

func (h *GetRoomByCodeHandler) GetRoomByCode(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	code := r.PathValue("code")
	// Получаем все комнаты через сервис
	rooms, err := h.GetRoomByCodeService.GetRoomByCode(ctx, code)
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
