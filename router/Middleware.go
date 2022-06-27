package router

import (
	"context"
	"github.com/horcrux12/clean-rest-api-template/config"
	"github.com/horcrux12/clean-rest-api-template/constanta"
	"github.com/horcrux12/clean-rest-api-template/exception"
	"github.com/horcrux12/clean-rest-api-template/helper"
	applicationModel2 "github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"net/http"
	"time"
)

func Middleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		corsOriginHandler(&responseWriter)
		responseWriter.Header().Set("Content-Type", "application/json")
		if request.Method == http.MethodOptions {
			return
		} else {
			var contextModel *applicationModel2.ContextModel
			defer func() {
				if r := recover(); r != nil {
					timestamp := time.Now()
					contextModel = request.Context().Value(constanta.ApplicationContextConstanta).(*applicationModel2.ContextModel)
					contextModel.LoggerModel.Time = int64(time.Since(timestamp).Seconds())
					exception.ErrorHandler(responseWriter, request, r, contextModel)
				}
			}()

			requestID := request.Header.Get(constanta.RequestIDConstanta)
			if requestID == "" {
				requestID = helper.GetUUID()
				request.Header.Set(constanta.RequestIDConstanta, requestID)
			}
			var contextModels applicationModel2.ContextModel

			contextModels.LoggerModel = applicationModel2.GenerateLogModel(config.ApplicationConfiguration.GetServerVersion(), config.ApplicationConfiguration.GetServerResourceApps())
			contextModels.LoggerModel.IP = request.Header.Get(constanta.IPAddressConstanta)
			contextModels.LoggerModel.Location = "[Middleware.go,Middleware]"

			ctx := context.WithValue(request.Context(), constanta.ApplicationContextConstanta, &contextModels)
			request = request.WithContext(ctx)

			nextHandler.ServeHTTP(responseWriter, request)

		}
	})
}

//func logMiddleware(loggerModel applicationModel.LoggerModel, requestURI string) {
//	if !strings.Contains(requestURI, "health") && !strings.Contains(requestURI, "docs") && !strings.Contains(requestURI, "swagger") {
//		util2.InputLog(errorModel.GenerateNonErrorModel(), loggerModel)
//	}
//}

func corsOriginHandler(responseWriter *http.ResponseWriter) {
	(*responseWriter).Header().Set("Access-Control-Allow-Origin", "*")
	(*responseWriter).Header().Set("Access-Control-Allow-Headers", "origin, content-type, accept, authorization")
	(*responseWriter).Header().Set("Access-Control-Allow-Credentials", "true")
	(*responseWriter).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
	(*responseWriter).Header().Set("Access-Control-Max-Age", "1209600")
}
