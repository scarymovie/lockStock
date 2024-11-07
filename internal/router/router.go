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
	roomRouter.HandleFunc("/find/{token}", roomHandler.FindRoomByToken)
	roomRouter.HandleFunc("/{roomId}/connect", roomHandler.ConnectToRoom)
	roomRouter.HandleFunc("/{roomId}/gamers/list", roomHandler.GetAllActiveRooms)

	mainRouter.Handle("/user/", http.StripPrefix("/user", userRouter))
	mainRouter.Handle("/room/", http.StripPrefix("/room", roomRouter))

	mainRouter.Handle("/api/", http.StripPrefix("/api", mainRouter))
}
