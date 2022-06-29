package CategoryService

import (
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
)

type CategoryService interface {
	Create(ctx *applicationModel.ContextModel, inputRequest in.CategoryRequest) (payload out.WebResponse)
	Update(ctx *applicationModel.ContextModel, inputRequest in.CategoryRequest) (payload out.WebResponse)
	Delete(ctx *applicationModel.ContextModel, inputRequest in.CategoryRequest) (payload out.WebResponse)
	FindByID(ctx *applicationModel.ContextModel, inputRequest in.CategoryRequest) (payload out.WebResponse)
	FindAll(ctx *applicationModel.ContextModel) (payload out.WebResponse)
}
