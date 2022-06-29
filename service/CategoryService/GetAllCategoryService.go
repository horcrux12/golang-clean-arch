package CategoryService

import (
	"github.com/horcrux12/clean-rest-api-template/app"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
)

func (service CategoryServiceImpl) FindAll(ctx *applicationModel.ContextModel) (payload out.WebResponse) {
	err := app.OpenTxConnection(ctx, app.ApplicationAttribute.DBConnection)
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(ctx.ConnectionModel.Tx)

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
