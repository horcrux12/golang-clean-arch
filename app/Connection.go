package app

import (
	"database/sql"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
)

func OpenTxConnection(ctx *applicationModel.ContextModel, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	ctx.Tx = tx
	return nil
}
