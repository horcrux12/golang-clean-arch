package exception

import (
	"github.com/go-playground/validator/v10"
	"github.com/horcrux12/clean-rest-api-template/config"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/model/errorModel"
	"net/http"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}, ctx *applicationModel.ContextModel) {
	switch err.(type) {
	case validator.ValidationErrors:
		validationErrors(writer, request, err, ctx)
		break
	case NotFoundError:
		notFoundError(writer, request, err, ctx)
		break
	case errorModel.ErrorExceptionModel:
		i18nError(writer, request, err, ctx)
		break
	default:
		internalServerError(writer, request, err, ctx)
		break
	}

	helper.LogError(ctx.LoggerModel.ToLoggerObject())
}

func validationErrors(writer http.ResponseWriter, request *http.Request, err interface{}, ctx *applicationModel.ContextModel) {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		var errDetail []string

		trans := errorModel.GetTranslatorForTranslate(ctx.Locale)
		for _, fieldError := range exception {
			errDetail = append(errDetail, fieldError.Translate(trans))
		}

		webResponse := out.PayloadStatusResponse{
			Code:    http.StatusBadRequest,
			Message: "BAD REQUEST",
			Detail:  errDetail,
		}

		ctx.LoggerModel.Message = helper.StructToJSON(webResponse.Detail)

		helper.WriteToResponseBody(writer, createResponseError(webResponse))
	} else {
		internalServerError(writer, request, err, ctx)
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}, ctx *applicationModel.ContextModel) {
	errException, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := out.PayloadStatusResponse{
			Code:    http.StatusNotFound,
			Message: "NOT FOUND",
			Detail:  errException.Error,
		}

		ctx.LoggerModel.Message = helper.StructToJSON(webResponse.Detail)

		helper.WriteToResponseBody(writer, createResponseError(webResponse))
	} else {
		internalServerError(writer, request, err, ctx)
	}
}

func i18nError(writer http.ResponseWriter, request *http.Request, err interface{}, ctx *applicationModel.ContextModel) {
	exception, ok := err.(errorModel.ErrorExceptionModel)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(exception.Code)

		webResponse := out.PayloadStatusResponse{
			Code:    exception.Code,
			Message: GenerateErrorI18NMessage(exception, helper.SwitchLanguageI18N(*ctx)),
			Detail:  exception.ErrorCause,
		}

		helper.CreateErrorLocation(exception.FileName, exception.FuncName, ctx)
		if exception.ErrorCause != nil {
			ctx.LoggerModel.Message = exception.ErrorCause.Error()
		} else {
			ctx.LoggerModel.Message = webResponse.Message
		}

		helper.WriteToResponseBody(writer, createResponseError(webResponse))
	} else {
		internalServerError(writer, request, err, ctx)
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}, ctx *applicationModel.ContextModel) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	webResponse := out.PayloadStatusResponse{
		Code:    http.StatusInternalServerError,
		Message: "INTERNAL SERVER ERROR",
		Detail:  err,
	}

	errDetail, ok := err.(error)
	if ok {
		ctx.LoggerModel.Message = errDetail.Error()
	} else {
		ctx.LoggerModel.Message = helper.StructToJSON(err)
	}
	helper.WriteToResponseBody(writer, createResponseError(webResponse))
}

func createResponseError(responseStatus out.PayloadStatusResponse) out.WebResponse {
	return out.WebResponse{
		Header: out.HeaderWebResponse{
			Version:   config.ApplicationConfiguration.GetServerVersion(),
			Timestamp: helper.GetTimeStamp(),
		},
		Payload: out.PayloadWebResponse{
			Status: responseStatus,
			Data:   nil,
			Other:  nil,
		},
	}
}
