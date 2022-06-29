package repository

import (
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
)

type UserRepository interface {
	CreateUser(ctx *applicationModel.ContextModel, user entity.UserModel) entity.UserModel
	UpdateUser(ctx *applicationModel.ContextModel, user entity.UserModel) entity.UserModel
	DeleteUser(ctx *applicationModel.ContextModel, user entity.UserModel)
	DetailUser(ctx *applicationModel.ContextModel, user entity.UserModel) entity.UserModel
	LoginUser(ctx *applicationModel.ContextModel, user entity.UserModel) entity.UserModel
	GetListUser(ctx *applicationModel.ContextModel) []entity.UserModel
}
