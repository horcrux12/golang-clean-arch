package UserService

import (
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
)

type UserService interface {
	CreateUser(ctx *applicationModel.ContextModel, inputRequest in.UserRequest) (payload out.WebResponse)
	UpdateUser(ctx *applicationModel.ContextModel, inputRequest in.UserRequest) (payload out.WebResponse)
	DeleteUser(ctx *applicationModel.ContextModel, inputRequest in.UserRequest) (payload out.WebResponse)
	UserLogin(ctx *applicationModel.ContextModel, inputRequest in.UserLoginRequest) (payload out.WebResponse)
	GetListUser(ctx *applicationModel.ContextModel) (payload out.WebResponse)
	DetailUser(ctx *applicationModel.ContextModel, inputRequest in.UserRequest) (payload out.WebResponse)
}
