# Базовый образ для сборки
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go

# Финальный образ
FROM alpine:latest

WORKDIR /app

# Копируем бинарник и конфиги
COPY --from=builder /app/main .
COPY configs ./configs

# Указываем порт
EXPOSE 8000

# Команда для запуска
CMD ["./main"]