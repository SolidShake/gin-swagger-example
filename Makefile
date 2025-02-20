.PHONY: run build test swagger clean

# Имя бинарного файла
BINARY_NAME=api-server

build:
	go build -o $(BINARY_NAME) cmd/main.go

run:
	go run cmd/main.go

test:
	go test -v ./...

swagger:
	swag init -g cmd/main.go

clean:
	rm -f $(BINARY_NAME)
	rm -rf ./docs

# Установка зависимостей
deps:
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@latest

# Запуск всех необходимых действий для старта проекта
all: deps swagger build run 