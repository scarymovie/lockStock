package room

import "time"

type Room struct {
	id        int    // приватный ID для базы данных
	UID       string // публичный UID
	Name      string
	Code      string
	CreatedAt time.Time // дата создания
}
