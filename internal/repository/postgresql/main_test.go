package postgresql

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"

)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://fituser:thisissupersecret@localhost:5432/fit_db?sslmode=disable"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect do db:", err)
	}

	testQueries = New(conn)

	m.Run()
}