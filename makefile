MIGRATE_DB_URL=mysql://root:12345@tcp(localhost:3306)/test

mysql:
	docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=12345 -e MYSQL_DATABASE=test -d mysql:8.0
	
migrate_up:
	migrate -path internal/db/migration -database "${MIGRATE_DB_URL}" -verbose up

migrate_down:
	migrate -path internal/db/migration -database "${MIGRATE_DB_URL}" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

docker_build:
	docker build -t test-server:latest .

proto:
	rm -rf internal/pb/*.go
	protoc --proto_path=internal/proto \
	--go_out=internal/pb --go_opt=paths=source_relative \
	--go-grpc_out=internal/pb --go-grpc_opt=paths=source_relative \
	internal/proto/*.proto

evans:
	evans --host localhost --port 8081 -r repl

.PHONY: mysql migrate_up migrate_down sqlc test server docker_build proto evans
