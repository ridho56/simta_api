package helper

import (
	"context"
	"database/sql"
	"time"
)

func GetConnection2() *sql.DB {
	db, err := sql.Open("mysql", "myfin:Admin@myfin123@tcp(103.189.234.90:3306)/d3ti_psdku")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func GetMahasiswaname(id string) string {
	tx := GetConnection2()
	ctx := context.Background()
	script := "SELECT nama FROM mahasiswa WHERE id_mhs = ?"
	rows, err := tx.QueryContext(ctx, script, id)
	PanicIfError(err)
	defer rows.Close()

	var name string

	if rows.Next() {
		rows.Scan(&name)
		return name
	}
	return name
}

func GetStafname(id string) string {
	tx := GetConnection2()
	ctx := context.Background()
	script := "SELECT nama FROM staf WHERE id_staf = ?"
	rows, err := tx.QueryContext(ctx, script, id)
	PanicIfError(err)
	defer rows.Close()

	var nama string

	if rows.Next() {
		rows.Scan(&nama)
		return nama
	}
	return nama
}
