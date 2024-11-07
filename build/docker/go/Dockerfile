FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY ./cmd/app /app
COPY go.mod go.sum ./
RUN go mod download
RUN go build -o app main.go

FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/app .
EXPOSE 8080
CMD ["./app"]
