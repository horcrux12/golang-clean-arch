package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/model/applicationModel"
	"github.com/horcrux12/clean-rest-api-template/repository"
	"github.com/horcrux12/clean-rest-api-template/service/UserService"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService UserService.UserService
	AbstractController
}

func NewUserController(db *sql.DB, validate *validator.Validate) UserController {
	userRepository := repository.NewUserRepository()
	return UserController{
		UserService: UserService.NewUserService(userRepository, db, validate),
	}
}

func (controller UserController) readBodyAndParam(request *http.Request) (result in.UserRequest) {
	strBody := helper.ReadBody(request)
	if strBody != "" {
		errorS := json.Unmarshal([]byte(strBody), &result)
		helper.PanicIfError(errorS)
	}

	id, _ := strconv.Atoi(mux.Vars(request)["ID"])
	if result.ID == 0 {
		result.ID = int64(id)
	}

	return
}

func (controller UserController) UserControllerWithoutPathParam(writer http.ResponseWriter, request *http.Request) {
	funcName := "UserControllerWithoutPathParam"
	var contextModel *applicationModel.ContextModel
	var payload out.WebResponse
	var inputRequest in.UserRequest

	inputRequest = controller.readBodyAndParam(request)

	switch request.Method {
	case http.MethodPost:
		contextModel = controller.WhiteListServe(funcName, writer, request)
		payload = controller.UserService.CreateUser(contextModel, inputRequest)
		break
	}

	defer controller.LogResponse(contextModel, payload, funcName)
	helper.WriteToResponseBody(writer, payload)
}
