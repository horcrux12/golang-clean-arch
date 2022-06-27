package controller

import (
	"context"
	"fmt"
	"github.com/horcrux12/clean-rest-api-template/constanta"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"net/http"
)

type AbstractController struct {
	FileName string
}

func (controller AbstractController) ServeController(funcName string, response http.ResponseWriter, request *http.Request, serveFunction func(*applicationModel.ContextModel, *http.Request) out.WebResponse) {
	contextModel := request.Context().Value(constanta.ApplicationContextConstanta).(*applicationModel.ContextModel)
	contextModel.Permission = "test"
	contextModel.AuthAccessModel.Locale = constanta.IDLangConstanta

	defer func() {
		contextModel.LoggerModel.Location = fmt.Sprintf("[%s, %s]", controller.FileName, funcName)
		helper.LogInfo(contextModel.LoggerModel.ToLoggerObject())
	}()

	ctx := context.WithValue(request.Context(), constanta.ApplicationContextConstanta, contextModel)
	request = request.WithContext(ctx)

	//realContext := contextModel.Value(constanta.ApplicationContextConstanta).(*applicationModel.ContextModel)
	payload := serveFunction(contextModel, request)
	contextModel.LoggerModel.Status = payload.Payload.Status.Code
	contextModel.LoggerModel.Message = payload.Payload.Status.Message

	helper.WriteToResponseBody(response, payload)
}

func ServeControllerClean[T any](funcName string, request *http.Request) *applicationModel.ContextModel {
	contextModel := request.Context().Value(constanta.ApplicationContextConstanta).(*applicationModel.ContextModel)
	contextModel.Permission = "test"
	contextModel.AuthAccessModel.Locale = constanta.IDLangConstanta

	defer func() {
		contextModel.LoggerModel.Location = fmt.Sprintf("[%s, %s]", "controller.FileName", funcName)
		helper.LogInfo(contextModel.LoggerModel.ToLoggerObject())
	}()

	ctx := context.WithValue(request.Context(), constanta.ApplicationContextConstanta, contextModel)
	request = request.WithContext(ctx)

	return contextModel
}
