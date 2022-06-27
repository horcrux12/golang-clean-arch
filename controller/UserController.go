package controller

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/horcrux12/clean-rest-api-template/repository"
	"github.com/horcrux12/clean-rest-api-template/service/UserService"
	"net/http"
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

func (controller UserController) UserControllerWithoutPathParam(writer http.ResponseWriter, request *http.Request) {
	funcName := "UserControllerWithoutPathParam"
	switch request.Method {
	case http.MethodPost:
		controller.ServeController(funcName, writer, request, controller.UserService.CreateUser)
		break
	}
}
