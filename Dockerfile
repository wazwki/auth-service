# Используем официальный образ Golang в качестве базового
FROM golang:1.22

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./

# Скачиваем зависимости
RUN go mod download

# Копируем исходный код в контейнер
COPY . .

# Сборка приложения
RUN go build -o auth-service ./cmd/main.go

# Открываем порт, который будет использоваться нашим приложением
EXPOSE 50051

# Запуск приложения
CMD ["./auth-service"]
