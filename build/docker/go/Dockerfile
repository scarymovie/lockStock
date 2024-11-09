# Первый этап: сборка приложения
FROM golang:1.23-alpine AS builder

# Устанавливаем рабочую директорию для сборки
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь код проекта
COPY . .

# Переходим в директорию с main.go
WORKDIR /app/cmd/app

# Сборка приложения
RUN go build -o app main.go

# Второй этап: создание минимального образа для запуска
FROM alpine:3.18

# Устанавливаем рабочую директорию для запуска
WORKDIR /app

# Копируем скомпилированное приложение из предыдущего этапа
COPY --from=builder /app/cmd/app/app .

# Открываем порт
EXPOSE 8080

# Команда запуска приложения
CMD ["./app"]
