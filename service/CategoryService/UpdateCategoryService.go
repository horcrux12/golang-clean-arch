package CategoryService

import (
	"database/sql"
	"github.com/horcrux12/clean-rest-api-template/app"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/exception"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
)

func (service CategoryServiceImpl) Update(ctx *applicationModel.ContextModel, inputRequest in.CategoryRequest) (payload out.WebResponse) {
	err := service.Validate.Struct(inputRequest)
	helper.PanicIfError(err)

	//tx, errDB := service.DB.Begin()
	//helper.PanicIfError(errDB)
	//defer helper.CommitOrRollback(tx)

	err = app.OpenTxConnection(ctx, app.ApplicationAttribute.DBConnection)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(ctx.ConnectionModel.Tx)

	categoryModel := entity.CategoryModel{
		Name: sql.NullString{String: inputRequest.Name},
		ID:   sql.NullInt64{Int64: inputRequest.ID},
	}

	// Update Data
	categoryModel, err = service.CategoryRepository.FindByID(ctx, entity.CategoryModel{ID: sql.NullInt64{Int64: inputRequest.ID}})
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	categoryModel.Name.String = inputRequest.Name

	categoryModel = service.CategoryRepository.Update(ctx, categoryModel)
	payload.Payload.Status = out.PayloadStatusResponse{
		Code:    200,
		Message: "OK",
	}

	output := service.ToCategoryResponse(categoryModel)
	payload.Payload.Data = output
	return
}
