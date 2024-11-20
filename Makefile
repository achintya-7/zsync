new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database "sqlite://./db/sqlite3.db" -verbose up

migratedown:
	migrate -path db/migration -database "sqlite://./db/sqlite3.db" -verbose down