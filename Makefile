createdb:
	docker exec -it postgres12 createdb --username=younes --owner=younes simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=younes -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

migrateup:
	migrate -path db/migration -database "postgres://younes:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://younes:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
