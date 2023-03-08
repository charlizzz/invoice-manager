postgres:
	docker run --name postgres-db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:alpine

createdb:
	docker exec -it postgres-db createdb --username=root --owner=root invoice-db

dropdb:
	docker exec -it postgres-db dropdb invoice-db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/invoice-db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/invoice-db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc