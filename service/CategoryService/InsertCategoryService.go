package CategoryService

import (
	"database/sql"
	"github.com/horcrux12/clean-rest-api-template/app"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
)

func (service CategoryServiceImpl) Create(ctx *applicationModel.ContextModel, inputRequest in.CategoryRequest) (payload out.WebResponse) {
	err := service.Validate.Struct(inputRequest)
	helper.PanicIfError(err)

	categoryModel := entity.CategoryModel{
		Name: sql.NullString{String: inputRequest.Name},
	}

	app.OpenTxConnection(ctx)
	defer helper.CommitOrRollback(ctx.ConnectionModel.Tx)

	// Insert Data
	categoryModel = service.CategoryRepository.Save(ctx, categoryModel)
	payload.Payload.Status = service.GetCommonResponseMessage("SUCCESS_INSERT_MESSAGE", ctx)

	output := service.ToCategoryResponse(categoryModel)
	payload.Payload.Data = output

	return
}

func (service CategoryServiceImpl) ToCategoryResponse(model entity.CategoryModel) out.CategoryResponse {
	return out.CategoryResponse{
		ID:   model.ID.Int64,
		Name: model.Name.String,
	}
}
