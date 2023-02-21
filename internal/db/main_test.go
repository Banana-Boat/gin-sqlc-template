package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbSource = "root:12345@tcp(localhost:3306)/test?parseTime=true"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can't connect to db", err)
	}

	testQueries = New(conn)
	fmt.Println("connect success")

	os.Exit(m.Run())
}
