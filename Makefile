build:
	docker-compose build houser

run:
	docker-compose up houser

test:
	go test -v ./...

postgres:
    docker run --name houser_db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
    docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
    docker exec -it postgres12 dropdb simple_bank

migrateup:
    migrate -path ./schema -database 'postgres://postgres:123123@0.0.0.0:5436/houser_db?sslmode=disable' up

migratedown:
    migrate -path ./schema -database 'postgres://postgres:123123@0.0.0.0:5436/houser_db?sslmode=disable' down

swag:
	swag init -g main.go
