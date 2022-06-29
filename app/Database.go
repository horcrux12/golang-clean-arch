package app

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"os"
)

func NewDB(address, defaultSchema string, maxOpenConn, maxIdleConn int) *sql.DB {
	dataSource := fmt.Sprintf("%s search_path=%s", address, defaultSchema)
	db, err := sql.Open("pgx", dataSource)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)

	return db
}
