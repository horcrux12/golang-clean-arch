package CategoryService

import (
	"database/sql"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/exception"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
	"net/http"
)

func (service CategoryServiceImpl) Update(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse) {
	var inputStruct in.CategoryUpdateRequest
	helper.ReadFromRequestBody(request, &inputStruct)

	err := service.Validate.Struct(inputStruct)
	helper.PanicIfError(err)

	categoryModel := entity.CategoryModel{
		Name: sql.NullString{String: inputStruct.Name},
		ID:   sql.NullInt64{Int64: inputStruct.ID},
	}

	// Update Data
	categoryModel, err = service.CategoryRepository.FindByID(ctx, entity.CategoryModel{ID: sql.NullInt64{Int64: inputStruct.ID}})
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	categoryModel.Name.String = inputStruct.Name

	categoryModel = service.CategoryRepository.Update(ctx, categoryModel)
	payload.Payload.Status = out.PayloadStatusResponse{
		Code:    200,
		Message: "OK",
	}

	output := service.ToCategoryResponse(categoryModel)
	payload.Payload.Data = output
	return
}
