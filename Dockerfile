# Этап 1: Сборка приложения
FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Собираем бинарник в бинарный файл "main"
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Этап 2: Финальный легковесный образ
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
# Копируем скомпилированный бинарник из предыдущего этапа
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
# Если есть статические файлы или .env, их тоже нужно скопировать
# COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./main"]