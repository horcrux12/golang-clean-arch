package CategoryService

import (
	"database/sql"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
	"github.com/horcrux12/clean-rest-api-template/model/errorModel"
	"net/http"
)

func (service CategoryServiceImpl) FindByID(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse) {
	funcName := "FindByID"
	inputStruct := service.ReadBodyAndParam(request)

	defer helper.CommitOrRollback(service.TX)

	category, err := service.CategoryRepository.FindByID(ctx, entity.CategoryModel{ID: sql.NullInt64{Int64: inputStruct.ID}})
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
