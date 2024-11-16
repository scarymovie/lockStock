// router.go
package router

import (
	"database/sql"
	handlers "lockStock/internal/network/http"
	"net/http"
)

// LoadRoutes загружает маршруты и передает подключение к базе данных в обработчики
func LoadRoutes(mainRouter *http.ServeMux, db *sql.DB) {
	// Под-маршрутизаторы для обработки пользователей и комнат
	userRouter := http.NewServeMux()
	roomRouter := http.NewServeMux()

	// Инициализация обработчиков с передачей `db`
	userHandler := handlers.NewUserHandler(db)
	getAllRoomsHandler := handlers.NewGetAllRoomsHandler(db)
	getRoomByCodeHandler := handlers.NewGetRoomByCodeHandler(db)

	// Регистрация маршрутов пользователя
	userRouter.HandleFunc("/create", userHandler.CreateUser)

	// Регистрация маршрутов комнат
	roomRouter.HandleFunc("/list", getAllRoomsHandler.GetAllActiveRooms)
	roomRouter.HandleFunc("/find/token/{code}", getRoomByCodeHandler.GetRoomByCode) // Без параметра для примера
	//roomRouter.HandleFunc("/connect/", getAllRoomsHandler.ConnectToRoom)       // Обработка id в обработчике
	//roomRouter.HandleFunc("/gamers/list", getAllRoomsHandler.GetAllActiveRooms)

	// Подключение под-маршрутизаторов
	mainRouter.Handle("/user/", http.StripPrefix("/user", userRouter))
	mainRouter.Handle("/room/", http.StripPrefix("/room", roomRouter))

	// Основной API путь
	mainRouter.Handle("/api/", http.StripPrefix("/api", mainRouter))
}
