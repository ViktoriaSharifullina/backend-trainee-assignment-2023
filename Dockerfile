# Используем официальный образ Golang как базовый
FROM golang:1.21

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Установка клиента PostgreSQL
RUN apt-get update && apt-get install -y postgresql-client

# Копируем файлы go.mod и go.sum и выполняем go get для зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы проекта в рабочую директорию
COPY . .

# Копируем SQL-скрипт в контейнер
COPY script.sql /scripts/

# Собираем приложение
RUN go build -o main .

CMD ["./main"]
