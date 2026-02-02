# Этап сборки
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Устанавливаем зависимости для сборки
RUN apk add --no-cache git

# Копируем модули и скачиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем код
COPY . .

RUN go mod tidy

# Собираем бинарник (включая миграции)
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server cmd/server/main.go

# Этап выполнения (минимальный образ)
FROM alpine:latest

WORKDIR /root/

# Устанавливаем сертификаты для HTTPS
RUN apk --no-cache add ca-certificates

# Копируем бинарник из этапа сборки
COPY --from=builder /app/server .

# Открываем порт
EXPOSE 3000

# Запускаем
CMD ["./server"]
