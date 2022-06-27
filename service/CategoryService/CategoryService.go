package CategoryService

import (
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"net/http"
)

type CategoryService interface {
	Create(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse)
	Update(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse)
	Delete(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse)
	FindByID(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse)
	FindAll(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse)
}
