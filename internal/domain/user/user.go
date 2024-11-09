package user

import (
	"time"
)

type User struct {
	id        int       // приватный ID для базы данных
	UID       string    // публичный UID
	CreatedAt time.Time // дата создания
}

func NewUser(uid string) *User {
	return &User{
		UID:       uid,
		CreatedAt: time.Now(),
	}
}
