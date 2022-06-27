package router

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/horcrux12/clean-rest-api-template/config"
	"net/http"
	"strconv"
)

func APIRouter() {
	InitiateRouter()

	handler := mux.NewRouter()
	prefixPath := config.ApplicationConfiguration.GetServerPrefixPath()
	if prefixPath != "" {
		prefixPath = "/" + prefixPath
	}

	handler.HandleFunc("/v1"+prefixPath+"/categories", AppRouter.CategoryController.CategoryControllerWithoutPathParam).Methods(http.MethodPost, http.MethodGet, http.MethodOptions)
	handler.HandleFunc("/v1"+prefixPath+"/categories/{ID}", AppRouter.CategoryController.CategoryControllerWithPathParam).Methods(http.MethodPut, http.MethodGet, http.MethodDelete, http.MethodOptions)

	handler.HandleFunc("/v1"+prefixPath+"/user", AppRouter.UserController.UserControllerWithoutPathParam).Methods(http.MethodPost, http.MethodGet, http.MethodOptions)

	//handler.HandleFunc("/v1"+prefixPath+"/get-ip", AppRouter.TrialController.TrialGetIPController).Methods(http.MethodGet, http.MethodOptions)

	handler.Use(Middleware)
	endpoint := config.ApplicationConfiguration.GetServerHost() + ":" + strconv.Itoa(config.ApplicationConfiguration.GetServerPort())
	fmt.Print(http.ListenAndServe(endpoint, handler))
}
