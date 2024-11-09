package router

import (
	handlers "lockStock/internal/network/http"
	"net/http"
)

func LoadRoutes(mainRouter *http.ServeMux) {
	userRouter := http.NewServeMux()
	roomRouter := http.NewServeMux()

	userHandler := &handlers.UserHandler{}
	roomHandler := &handlers.RoomHandler{}

	userRouter.HandleFunc("/create", userHandler.CreateUser)

	roomRouter.HandleFunc("/list", roomHandler.GetAllActiveRooms)
	roomRouter.HandleFunc("/find/token/{token}", roomHandler.FindRoomByToken)
	roomRouter.HandleFunc("/connect/{roomId}/", roomHandler.ConnectToRoom)
	roomRouter.HandleFunc("/gamers/{roomId}/list", roomHandler.GetAllActiveRooms)

	mainRouter.Handle("/user/", http.StripPrefix("/user", userRouter))
	mainRouter.Handle("/room/", http.StripPrefix("/room", roomRouter))

	mainRouter.Handle("/api/", http.StripPrefix("/api", mainRouter))
}
