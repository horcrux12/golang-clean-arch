package CategoryService

import (
	"database/sql"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
	"net/http"
)

func (service CategoryServiceImpl) Create(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse) {
	var inputStruct in.CategoryCreateRequest
	helper.ReadFromRequestBody(request, &inputStruct)

	err := service.Validate.Struct(inputStruct)
	helper.PanicIfError(err)

	categoryModel := entity.CategoryModel{
		Name: sql.NullString{String: inputStruct.Name},
	}

	// Insert Data
	categoryModel = service.CategoryRepository.Save(ctx, categoryModel)
	payload.Payload.Status = out.PayloadStatusResponse{
		Code:    200,
		Message: "OK",
	}

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
