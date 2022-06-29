package CategoryService

import (
	"database/sql"
	"github.com/horcrux12/clean-rest-api-template/app"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
	"github.com/horcrux12/clean-rest-api-template/model/errorModel"
)

func (service CategoryServiceImpl) FindByID(ctx *applicationModel.ContextModel, inputRequest in.CategoryRequest) (payload out.WebResponse) {
	funcName := "FindByID"
	validateErr := service.Validate.Var(inputRequest.ID, "required")
	helper.PanicIfErrorWithLocation(validateErr, service.FileName, funcName, ctx)

	err := app.OpenTxConnection(ctx, app.ApplicationAttribute.DBConnection)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(ctx.ConnectionModel.Tx)

	category, err := service.CategoryRepository.FindByID(ctx, entity.CategoryModel{ID: sql.NullInt64{Int64: inputRequest.ID}})
	if err != nil {
		panic(errorModel.GenerateDataNotFound(service.FileName, funcName))
	}
	payload.Payload.Status = out.PayloadStatusResponse{
		Code:    200,
		Message: "OK",
	}

	output := service.ToCategoryResponse(category)
	payload.Payload.Data = output
	return
}
