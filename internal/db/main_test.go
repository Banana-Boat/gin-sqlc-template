package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbDriver = "mysql"
	dbSource = "mysql://root:12345@tcp(localhost:3307)/test"
)

var testQuery *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("can't connect to db", err)
	}

	testQuery = New(conn)

	os.Exit(m.Run())
}
