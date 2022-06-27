package CategoryService

import (
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
	"net/http"
)

func (service CategoryServiceImpl) FindAll(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse) {
	defer helper.CommitOrRollback(service.TX)

	categoriesOnDB := service.CategoryRepository.FindAll(ctx)
	payload.Payload.Status = service.GetCommonResponseMessage("SUCCESS_GET_LIST_MESSAGE", ctx)

	payload.Payload.Data = service.ToCategoriesResponses(categoriesOnDB)
	return
}

func (service CategoryServiceImpl) ToCategoriesResponses(model []entity.CategoryModel) (result []out.CategoryResponse) {
	for _, categoryModel := range model {
		result = append(result, service.ToCategoryResponse(categoryModel))
	}
	return
}
