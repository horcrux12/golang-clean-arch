package CategoryService

import (
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"net/http"
)

func (service CategoryServiceImpl) Delete(ctx *applicationModel.ContextModel, request *http.Request) (payload out.WebResponse) {
	panic("implement me")
}
