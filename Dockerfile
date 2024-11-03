# Используем официальный образ Go для сборки
FROM golang:1.22.5-alpine AS builder

# Устанавливаем необходимые зависимости
RUN apk update && apk add --no-cache git

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем весь проект
COPY . .

# Собираем приложение
RUN go build -o main ./cmd/cyrs/main.go

# Используем минимальный образ для запуска
FROM alpine:latest

# Устанавливаем необходимые библиотеки
RUN apk --no-cache add ca-certificates

# Устанавливаем рабочую директорию
WORKDIR /root/

# Копируем бинарник из builder
COPY --from=builder /app/main .

# Копируем конфигурационный файл
COPY config/local.yaml ./config/local.yaml

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]
