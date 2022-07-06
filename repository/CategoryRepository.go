package repository

import (
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
)

type CategoryRepository interface {
	Save(ctx *applicationModel.ContextModel, category entity.CategoryModel) entity.CategoryModel
	Update(ctx *applicationModel.ContextModel, category entity.CategoryModel) entity.CategoryModel
	Delete(ctx *applicationModel.ContextModel, category entity.CategoryModel)
	FindByID(ctx *applicationModel.ContextModel, category entity.CategoryModel) (entity.CategoryModel, error)
	FindAll(ctx *applicationModel.ContextModel) []entity.CategoryModel
}
