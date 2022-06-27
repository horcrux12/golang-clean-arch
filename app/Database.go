package app

import (
	"database/sql"
	"fmt"
	"github.com/horcrux12/clean-rest-api-template/helper"
	_ "github.com/jackc/pgx/stdlib"
)

func NewDB(address, defaultSchema string, maxOpenConn, maxIdleConn int) *sql.DB {
	dataSource := fmt.Sprintf("%s search_path=%s", address, defaultSchema)
	db, err := sql.Open("pgx", dataSource)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)

	return db
}
