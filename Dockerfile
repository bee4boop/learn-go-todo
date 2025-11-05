# Используем официальный Go образ
FROM golang:1.24-alpine

# Создаём рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum и качаем зависимости
COPY go.mod go.sum ./
RUN go mod tidy

# Копируем весь код
COPY . .

# Сборка бинарника
RUN go build -o server main.go

# Запуск приложения
CMD ["./server"]
