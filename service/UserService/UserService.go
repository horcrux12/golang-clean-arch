package UserService

import (
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"net/http"
)

type UserService interface {
	CreateUser(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse)
	UpdateUser(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse)
	DeleteUser(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse)
	UserLogin(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse)
	GetListUser(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse)
	DetailUser(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse)
}
