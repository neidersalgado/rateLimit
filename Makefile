# Nombre de tu imagen Docker
IMAGE_NAME := ratelimit

.PHONY: build run test clean

build:
	@echo "Construyendo la imagen de Docker (${IMAGE_NAME})..."
	ENV=dev docker-compose build

run:
	@echo "Iniciando servicios..."
	ENV=dev docker-compose up -d

down:
	@echo "Deteniendo servicios..."
	ENV=dev docker-compose down

clean:
	@echo "Limpiando..."
	docker system prune -a
