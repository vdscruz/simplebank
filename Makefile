postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate --path db/migrate --database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate --path db/migrate --database "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

start:
	go run main.go

.PHONY: createdb dropdb postgres up down sqlc migrateup migratedown test start