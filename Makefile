newMigration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database "sqlite://./db/sqlite3.db" -verbose up

migratedown:
	migrate -path db/migration -database "sqlite://./db/sqlite3.db" -verbose down

gen:
	sqlc generate

test:
	go test -v -cover ./...

build:
	go build -o bin/main main.go

init:
	go run main.go init

history:
	go run main.go history