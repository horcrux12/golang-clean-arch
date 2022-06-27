package controller

import (
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/repository"
	"github.com/horcrux12/clean-rest-api-template/service/CategoryService"
	"net/http"
)

type CategoryController struct {
	CategoryService CategoryService.CategoryService
	AbstractController
}

//var CategoryController = CategoryController{}.New()

func NewCategoryController(db *sql.DB, validate *validator.Validate) CategoryController {
	tx, err := db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	categoryRepository := repository.NewCategoryRepository(tx)
	return CategoryController{
		CategoryService: CategoryService.NewCategoryService(categoryRepository, validate),
		AbstractController: AbstractController{
			FileName: "CategoryController.go",
		},
	}
}

func (controller CategoryController) CategoryControllerWithoutPathParam(writer http.ResponseWriter, request *http.Request) {
	funcName := "CategoryControllerWithoutPathParam"
	switch request.Method {
	case "POST":
		controller.ServeController(funcName, writer, request, controller.CategoryService.Create)
		break
	case "GET":
		controller.ServeController(funcName, writer, request, controller.CategoryService.FindAll)
		break
	}
}

func (controller CategoryController) CategoryControllerWithPathParam(writer http.ResponseWriter, request *http.Request) {
	funcName := "CategoryControllerWithPathParam"
	switch request.Method {
	case http.MethodGet:
		controller.ServeController(funcName, writer, request, controller.CategoryService.FindByID)
		break
	case http.MethodPut:
		controller.ServeController(funcName, writer, request, controller.CategoryService.Update)
		break
	case http.MethodDelete:
		controller.ServeController(funcName, writer, request, controller.CategoryService.Delete)
	}
}
