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
	roomHandler := handlers.NewRoomHandler(db)

	// Регистрация маршрутов пользователя
	userRouter.HandleFunc("/create", userHandler.CreateUser)

	// Регистрация маршрутов комнат
	roomRouter.HandleFunc("/list", roomHandler.GetAllActiveRooms)
	//roomRouter.HandleFunc("/find/token/", roomHandler.FindRoomByToken) // Без параметра для примера
	//roomRouter.HandleFunc("/connect/", roomHandler.ConnectToRoom)       // Обработка id в обработчике
	//roomRouter.HandleFunc("/gamers/list", roomHandler.GetAllActiveRooms)

	// Подключение под-маршрутизаторов
	mainRouter.Handle("/user/", http.StripPrefix("/user", userRouter))
	mainRouter.Handle("/room/", http.StripPrefix("/room", roomRouter))

	// Основной API путь
	mainRouter.Handle("/api/", http.StripPrefix("/api", mainRouter))
}
