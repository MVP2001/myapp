# Этап сборки
FROM mirror.gcr.io/library/golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates tzdata

ENV GOPROXY=https://proxy.golang.org,direct
# Отключаем проверку go.sum временно
ENV GOSUMDB=off
ENV GOFLAGS="-mod=mod"

COPY . .

# Скачиваем все зависимости и собираем
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server cmd/server/main.go

# Этап выполнения
FROM alpine:latest

WORKDIR /root/

RUN apk --no-cache add ca-certificates tzdata

COPY --from=builder /app/server .

EXPOSE 3000

CMD ["./server"]
