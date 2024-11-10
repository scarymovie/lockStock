package http

import (
	"database/sql"
	"log"
	"net/http"
)

type RoomHandler struct {
	DB *sql.DB
}

func (h *RoomHandler) GetAllActiveRooms(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	log.Println("received request to get all active rooms")
	w.Write([]byte("room tested"))
}

func (h *RoomHandler) FindRoomByToken(w http.ResponseWriter, r *http.Request) {
	token := r.PathValue("token")
	w.WriteHeader(http.StatusAccepted)
	log.Println("received request to find room by token")
	w.Write([]byte("you send token :" + token))
}

func (h *RoomHandler) ConnectToRoom(w http.ResponseWriter, r *http.Request) {
	roomId := r.PathValue("roomId")
	w.WriteHeader(http.StatusAccepted)
	log.Println("received request to connect to the room")
	w.Write([]byte("you send roomId :" + roomId))
}

func (h *RoomHandler) GetRoomGamers(w http.ResponseWriter, r *http.Request) {
	roomId := r.PathValue("roomId")
	w.WriteHeader(http.StatusAccepted)
	log.Println("received request to get room gamers")
	w.Write([]byte("you send roomId :" + roomId))
}
