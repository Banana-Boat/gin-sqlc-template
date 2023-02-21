mysql:
	docker run --name mysql8 -p 3307:3306 -e MYSQL_ROOT_PASSWORD=12345 -e MYSQL_DATABASE=test -d mysql:8.0

migrateup:
	migrate -path internal/db/migration -database "mysql://root:12345@tcp(localhost:3307)/test" -verbose up

migratedown:
	migrate -path internal/db/migration -database "mysql://root:12345@tcp(localhost:3307)/test" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

.PHONY: mysql mysqlbash migrateup migratedown sqlc test server
