MYSQL_URL=mysql://root:12345@tcp(localhost:3306)/test

mysql:
	docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=12345 -e MYSQL_DATABASE=test -d mysql:8.0

mysql_rm:
	docker stop mysql && docker rm mysql

mysql_bash:
	docker exec -it mysql bash
	
migrate_up:
	migrate -path internal/db/migration -database "${MYSQL_URL}" -verbose up

migrate_down:
	migrate -path internal/db/migration -database "${MYSQL_URL}" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...

server:
	go run main.go

.PHONY: mysql mysql_rm mysql_bash migrate_up migrate_down sqlc test server
