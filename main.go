package main

import (
	"database/sql"
	"log"

	"github.com/Banana-Boat/gin-sqlc-template/internal/api"
	"github.com/Banana-Boat/gin-sqlc-template/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver      = "mysql"
	dbSource      = "root:12345@tcp(localhost:3306)/test?parseTime=true"
	serverAddress = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can't connect to db", err)
	}

	store := db.NewStore(conn)

	server, err := api.NewServer(store)
	if err != nil {
		log.Fatal("cannot create server")
	}

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
