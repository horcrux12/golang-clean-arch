package repository

import (
	"database/sql"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/entity"
)

type UserRepository interface {
	CreateUser(ctx *applicationModel.ContextModel, tx *sql.Tx, user entity.UserModel) entity.UserModel
	UpdateUser(ctx *applicationModel.ContextModel, tx *sql.Tx, user entity.UserModel) entity.UserModel
	DeleteUser(ctx *applicationModel.ContextModel, tx *sql.Tx, user entity.UserModel)
	DetailUser(ctx *applicationModel.ContextModel, tx *sql.Tx, user entity.UserModel) entity.UserModel
	LoginUser(ctx *applicationModel.ContextModel, tx *sql.Tx, user entity.UserModel) entity.UserModel
	GetListUser(ctx *applicationModel.ContextModel, tx *sql.Tx) []entity.UserModel
}
