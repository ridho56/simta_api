package app

import (
	"database/sql"
	"testing"
	"time"
)

func TestDatabase(t *testing.T) {

	db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/d3ti_psdku")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	db.Close()
}
