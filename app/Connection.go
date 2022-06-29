package app

import (
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
)

func OpenTxConnection(ctx *applicationModel.ContextModel) {
	tx, err := ApplicationAttribute.DBConnection.Begin()
	helper.PanicIfError(err)
	ctx.Tx = tx
}
